# DevOps Lab — Resumo da sessão

**Sessão:** noite de 13/10/2026 (horário de Mato Grosso; 14/10 em UTC)  
**Ambiente:** VM 113 `devops-lab` — Debian 12, 2 vCPU, 4 GB RAM e 40 GB de disco.  
**Etapa:** preparação do ambiente para Kubernetes com K3s. **K3s ainda não instalado.**

## 1. Arquitetura atual

```text
Computador Omarchy (WireGuard: 10.8.0.2)
       |
       | VPN WireGuard
       v
VM 101 — roteador WireGuard (10.0.0.1)
       |
       | rede privada vmbr1
       v
VM 113 — devops-lab
  - enp6s19: 10.0.0.14 (rede privada; SSH)
  - eth0: 173.208.136.198 (interface pública)
  - Debian 12
  - Futuro K3s single-node: Control Plane + Worker
```

A administração via SSH está configurada para escutar em `10.0.0.14:22`. Não fizemos alterações na VM 101, no host Proxmox nem nas outras VMs.

## 2. Conceitos estudados

### Kernel Linux

É a parte central do sistema operacional que gerencia recursos como CPU, memória, armazenamento e rede. Os containers e o Kubernetes dependem desses recursos.

### Módulos do kernel

São componentes que adicionam funcionalidades ao kernel:

- **`overlay`**: suporte ao OverlayFS, usado para organizar camadas de arquivos de imagens e containers.
- **`br_netfilter`**: permite que determinados pacotes atravessando bridges Linux passem por mecanismos de filtragem de rede.

O `modprobe` identifica os módulos e suas dependências. O parâmetro `-n` **simula** o carregamento, sem modificar o sistema; `-v` mostra as operações previstas.

### IP forwarding

O parâmetro `net.ipv4.ip_forward` determina se o Linux pode encaminhar pacotes IPv4 entre interfaces:

- `0`: desabilitado.
- `1`: habilitado.

Habilitar forwarding não abre automaticamente portas, mas, combinado com rotas e políticas permissivas, pode permitir trânsito entre redes.

### Firewall

As cadeias consultadas do `iptables` são:

- **INPUT**: tráfego destinado à VM.
- **OUTPUT**: tráfego originado na VM.
- **FORWARD**: tráfego encaminhado através da VM.

### LLMNR

Protocolo de resolução de nomes em redes locais, associado à porta 5355. Não é necessário para este laboratório; foi desabilitado na VM 113.

### Kubernetes, K3s e containerd

O Kubernetes coordena a execução e o gerenciamento de aplicações em containers. O K3s é uma distribuição leve de Kubernetes. No nosso plano, um único nó terá os papéis de **Control Plane** e **Worker**; o `containerd` será o runtime responsável pelo gerenciamento dos containers.

## 3. Comandos executados e resultados

### Verificação dos módulos carregados

```bash
lsmod | grep -E 'overlay|br_netfilter'
```

**Resultado:** nenhuma linha retornada. Isso não prova que os módulos estejam indisponíveis.

### Consulta do encaminhamento IPv4

```bash
/usr/sbin/sysctl net.ipv4.ip_forward
```

**Resultado inicial:** `net.ipv4.ip_forward = 0`.

Foi feito um teste temporário de ativação:

```bash
sudo /usr/sbin/sysctl -w net.ipv4.ip_forward=1
/usr/sbin/sysctl net.ipv4.ip_forward
```

**Resultado:** `1`. Por precaução, retornamos à condição inicial:

```bash
sudo /usr/sbin/sysctl -w net.ipv4.ip_forward=0
```

**Estado ao final:** forwarding IPv4 desabilitado (`0`). A alteração por `sysctl -w` não foi gravada como configuração persistente.

### Inspeção do firewall

```bash
sudo nft list ruleset
sudo iptables -S
sudo ip6tables -S
```

**Resultados:** `nft` não foi encontrado; nas cadeias padrão consultadas pelo `iptables` e `ip6tables`, `INPUT`, `FORWARD` e `OUTPUT` apresentaram política `ACCEPT`, sem regras adicionais na saída. Isso **não comprova** ausência de outras camadas de firewall (como Proxmox ou provedor).

### Portas em escuta

```bash
sudo ss -lntup
```

**Resultados relevantes antes da alteração:** SSH escutando em `10.0.0.14:22`; `systemd-resolved` escutando na porta 5355 (LLMNR), inclusive em endereços de todas as interfaces.

### Desativação do LLMNR

Primeiro foi verificada a configuração:

```bash
resolvectl status
```

Depois foi criado um arquivo de configuração **somente na VM 113**:

```bash
sudo mkdir -p /etc/systemd/resolved.conf.d
printf '[Resolve]\nLLMNR=no\n' | sudo tee /etc/systemd/resolved.conf.d/10-disable-llmnr.conf
sudo systemctl restart systemd-resolved
```

Conteúdo do arquivo:

```ini
[Resolve]
LLMNR=no
```

**Validação:** `resolvectl status` apresentou `-LLMNR` globalmente, sem escopos LLMNR ativos nas interfaces. Embora as linhas das interfaces ainda mostrassem `+LLMNR` como configuração local, a porta 5355 deixou de escutar.

```bash
sudo ss -lntup | grep 5355
getent hosts github.com
```

**Resultados:** nenhum socket encontrado na porta 5355; `github.com` resolveu para `140.82.113.3`, confirmando que o DNS convencional continuou funcionando.

### Disponibilidade dos módulos

```bash
sudo /usr/sbin/modprobe -n -v overlay
sudo /usr/sbin/modprobe -n -v br_netfilter
```

**Resultado:** a simulação mostrou os módulos disponíveis no kernel `6.1.0-53-cloud-amd64` e as operações previstas:

```text
insmod .../overlay.ko
insmod .../llc.ko
insmod .../stp.ko
insmod .../bridge.ko
insmod .../br_netfilter.ko
```

Os módulos **não foram carregados** por esses comandos: `-n` executa apenas uma simulação.

## 4. Estado ao encerrar

| Item | Situação |
|---|---|
| SSH pela rede privada | Configurado em `10.0.0.14:22` |
| DNS | Funcionando |
| LLMNR | Desabilitado; porta 5355 sem listener |
| IP forwarding IPv4 | Desabilitado (`0`) |
| Políticas padrão `iptables` / `ip6tables` | `ACCEPT` nas cadeias consultadas |
| Módulos `overlay` e `br_netfilter` | Disponíveis; carregamento ainda não realizado |
| Firewall específico para o laboratório | Pendente |
| K3s | Não instalado |

## 5. Próxima sessão

1. Planejar as regras de firewall da VM 113 sem perder acesso SSH pela VPN e sem prejudicar a futura rede dos Pods.
2. Verificar o carregamento efetivo dos módulos e definir se são necessários parâmetros persistentes do kernel.
3. Revisar a instalação segura do K3s single-node: Control Plane, Worker, `containerd`, API Kubernetes e portas utilizadas.
4. Avaliar a desativação inicial de Traefik e ServiceLB para evitar exposição pública acidental.
5. Antes de componentes mais pesados, reavaliar os recursos da VM e a capacidade do host Proxmox.

**Decisão de segurança:** não expor a API Kubernetes nem aplicações publicamente até configurar e testar as regras de acesso. Nenhuma mudança deve ser feita nas outras VMs sem autorização.

---

**Ideia-chave:** Kubernetes organiza aplicações em containers; o kernel Linux fornece os mecanismos fundamentais de processos, arquivos e rede que tornam sua execução possível.
