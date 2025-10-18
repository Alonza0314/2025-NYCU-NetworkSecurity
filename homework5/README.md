# Homework 5

> [!Note]
> The system is **Ubuntu 20.04**.

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

### 2.2.4 Ensure telnet client is not installed

**Analysis:**

**Why OS left default configurations:**

- **Service Reachability**: telnet was historically used for remote terminal access and network troubleshooting
- **User Friendliness**: Simple command-line tool for quick connectivity testing and remote administration
- **Ease-of-Use**: Lightweight client with minimal configuration requirements, commonly used in legacy systems

**Security Issues:**

- **Unencrypted transmission**: All data including credentials transmitted in plain text
- **Credential theft risk**: Passwords and sensitive information vulnerable to network sniffing
- **No authentication integrity**: No protection against man-in-the-middle attacks
- **Legacy protocol**: Designed in 1969 without modern security considerations

**Remediation Solutions:**

1. **Complete removal** (recommended):

    ```bash
    apt purge telnet
    ```

2. **Secure alternatives**:

    - Use SSH for encrypted remote access: `ssh user@hostname`
    - Use modern tools like `nc` (netcat) for network testing
    - Implement VPN solutions for secure remote connectivity

3. **Temporary usage** (if absolutely necessary):

    - Install only when needed: `apt install telnet`
    - Remove immediately after use: `apt purge telnet`
    - Use only in isolated, secure environments

**Impact on Users:**

- **Positive**: Eliminates credential theft risk, forces adoption of secure protocols, improves overall security posture
- **Negative**: Loss of quick troubleshooting tool, need to learn SSH commands, potential workflow disruption for legacy system administrators

**Security vs User Friendliness Balance:**

The optimal approach is **complete removal with secure alternatives**:

- **Immediate action**: Remove telnet client completely
- **User training**: Provide SSH training and documentation
- **Troubleshooting tools**: Offer secure alternatives (SSH, netcat, modern network utilities)
- **Legacy support**: Use SSH tunneling for accessing legacy systems

This approach prioritizes security while providing users with more robust and secure alternatives that offer better functionality than telnet.

## 3 Network

## 4 HostBased Firewall

## 5 Access Control

## Contribution Table

| Student ID | Works | Percentage |
| - | - | - |
| 314581015 | Chapter 1 | 20% |
| 313581047 | Chapter 2 | 20% |
| 313581038 | Chapter 3 | 20% |
| 313581055 | Chapter 4 | 20% |
| 412581005 | Chapter 5 | 20% |
