# Homework 5

## 1 Initial Setup

## 2 Services

### 2.1.13 Ensure rsync services are not in use

**Analysis:**

**Why OS left default configurations:**

- **Service Reachability**: rsync is commonly used for file synchronization, backup operations, and system maintenance across networks
- **User Friendliness**: Pre-installed for immediate use by system administrators and developers
- **Ease-of-Use**: Simple command-line interface, widely supported by automation scripts and tools

**Security Issues:**

- Unencrypted data transmission over network
- Weak authentication mechanisms by default
- Network exposure increases attack surface
- Potential for unauthorized file access and data leakage

**Remediation Solutions:**

1. **Complete removal** (if not needed):

    ```bash
    systemctl stop rsync.service
    apt purge rsync
    ```

2. **Service masking** (if package required as dependency):

    ```bash
    systemctl stop rsync.service
    systemctl mask rsync.service
    ```

3. **Secure alternatives**:

    - Use rsync over SSH: `rsync -avz -e "ssh" user@host:/path/ /local/path/`
    - Implement proper authentication and access controls
    - Use modern encrypted sync tools (Syncthing, Nextcloud)

**Impact on Users:**

- **Positive**: Enhanced security, reduced attack surface, compliance with security standards
- **Negative**: Loss of convenient file sync capability, need to modify existing scripts, learning curve for alternatives

**Security vs User Friendliness Balance:**

The optimal solution is a **layered approach**:

- Immediately disable unnecessary rsync services
- Migrate critical functions to secure alternatives (rsync over SSH)
- Provide automated scripts and documentation for users
- Implement gradual transition with user training

This approach maintains functionality while significantly improving security posture without major disruption to user workflows.

## 3 Network

## 4 HostBased Firewall

## 5 Access Control