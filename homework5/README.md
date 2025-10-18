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

### 2.2.6 Ensure ftp client is not installed

**Analysis:**

**Why OS left default configurations:**

- **Service Reachability**: FTP is a traditional standard for file transfers, widely supported across different systems and platforms
- **User Friendliness**: Simple command-line interface for quick file uploads/downloads, especially for anonymous access scenarios
- **Ease-of-Use**: Lightweight client with minimal configuration, commonly used in legacy environments and automated scripts

**Security Issues:**

- **Unencrypted data transmission**: All file contents and credentials sent in plain text
- **Credential exposure**: Usernames and passwords vulnerable to network interception
- **No data integrity protection**: No verification that files haven't been tampered with during transfer
- **Anonymous access risk**: Many FTP servers allow anonymous connections without authentication
- **Legacy protocol vulnerabilities**: Designed in 1971 without modern security considerations

**Remediation Solutions:**

1. **Complete removal** (recommended):

    ```bash
    apt purge ftp
    ```

2. **Secure alternatives**:
   - Use SFTP for encrypted file transfers: `sftp user@hostname`
   - Use SCP for secure file copying: `scp file user@host:/path/`
   - Use modern tools like `rsync` over SSH: `rsync -avz -e ssh user@host:/path/ /local/`
   - Implement cloud storage solutions with encryption

3. **Temporary usage** (if absolutely necessary):
   - Install only when needed: `apt install ftp`
   - Remove immediately after use: `apt purge ftp`
   - Use only in isolated, secure networks

**Impact on Users:**

- **Positive**: Eliminates credential theft risk, forces adoption of secure file transfer methods, improves data protection
- **Negative**: Loss of simple file transfer tool, need to learn SFTP/SCP commands, potential disruption to legacy workflows

**Security vs User Friendliness Balance:**

The optimal approach is **complete removal with secure alternatives**:

- **Immediate action**: Remove FTP client completely
- **User training**: Provide SFTP/SCP training and documentation
- **Migration support**: Help users transition to secure alternatives
- **Legacy compatibility**: Use SSH tunneling for accessing legacy FTP servers

This approach prioritizes data security while providing users with more robust and encrypted file transfer capabilities that offer better protection than traditional FTP.

### 2.3.2.1 Ensure systemd-timesyncd configured with authorized timeserver

**Analysis:**

**Why OS left default configurations:**

- **Service Reachability**: Default NTP servers provide basic time synchronization without requiring manual configuration
- **User Friendliness**: Automatic time synchronization works out-of-the-box, reducing setup complexity for users
- **Ease-of-Use**: No additional configuration needed, system automatically connects to public NTP servers
- **Universal compatibility**: Default servers are widely accessible and generally reliable

**Security Issues:**

- **Untrusted time sources**: Default NTP servers may not be from authorized or trusted sources
- **Time manipulation attacks**: Malicious NTP servers could provide incorrect time, affecting security mechanisms
- **Log integrity concerns**: Inconsistent timestamps across systems complicate forensic investigations
- **Compliance violations**: Many security standards require time synchronization with authorized servers
- **Kerberos authentication failures**: Incorrect time can cause Kerberos ticket validation to fail

**Remediation Solutions:**

1. **Configure authorized NTP servers**:

    ```bash
    # Create configuration file
    sudo mkdir -p /etc/systemd/timesyncd.conf.d/
    
    # Add authorized time servers
    cat > /etc/systemd/timesyncd.conf.d/60-timesyncd.conf << EOF
    [Time]
    NTP=time.nist.gov
    FallbackNTP=time-a-g.nist.gov time-b-g.nist.gov time-c-g.nist.gov
    EOF
    
    # Reload configuration
    systemctl reload-or-restart systemd-timesyncd
    ```

2. **Verify configuration**:

    ```bash
    # Check current configuration
    timedatectl show-timesync --all
    
    # Verify time synchronization status
    timedatectl status
    ```

3. **Alternative secure time sources**:

    - Use internal NTP servers for enterprise environments
    - Configure GPS-based time sources for critical systems
    - Implement redundant time sources for high availability

**Impact on Users:**

- **Positive**: Improved security posture, consistent logging across systems, compliance with security standards
- **Negative**: Requires initial configuration, potential connectivity issues if authorized servers are unreachable
- **Minimal disruption**: Once configured, operates transparently without user intervention

**Security vs User Friendliness Balance:**

The optimal approach is **configured with fallback options**:

- **Primary action**: Configure authorized NTP servers with known, trusted sources
- **Fallback strategy**: Provide multiple fallback servers to ensure availability
- **Monitoring**: Implement time drift monitoring and alerting
- **Documentation**: Provide clear configuration examples and troubleshooting guides

This approach ensures reliable time synchronization while maintaining security through the use of authorized time sources, with minimal impact on user experience once properly configured.

### 2.4.1.2 Ensure access to /etc/crontab is configured

**Analysis:**

**Why OS left default configurations:**

- **Service Reachability**: Default permissions allow system administrators to easily manage cron jobs without complex permission changes
- **User Friendliness**: Permissive access enables quick troubleshooting and job management by authorized users
- **Ease-of-Use**: Standard file permissions reduce complexity for system maintenance and automation scripts
- **Legacy compatibility**: Historical Unix systems used more permissive defaults for system files

**Security Issues:**

- **Privilege escalation risk**: Unauthorized write access could allow users to execute commands as root
- **Information disclosure**: Read access reveals system job schedules and commands, aiding attackers in reconnaissance
- **Cron job manipulation**: Malicious users could modify scheduled tasks to gain persistent access
- **System integrity compromise**: Unauthorized modifications could lead to system compromise or data exfiltration
- **Compliance violations**: Many security standards require strict access controls on system configuration files

**Remediation Solutions:**

1. **Set proper ownership and permissions**:

    ```bash
    # Set root ownership
    chown root:root /etc/crontab
    
    # Remove group and other access, keep only owner (root) access
    chmod og-rwx /etc/crontab
    
    # Verify permissions (should show -rw-------)
    ls -l /etc/crontab
    ```

2. **Verify configuration**:

    ```bash
    # Check file ownership
    stat -c "%U:%G %a" /etc/crontab
    
    # Verify only root can read/write
    sudo -u nobody cat /etc/crontab  # Should fail
    ```

3. **Additional security measures**:

    - Implement file integrity monitoring (AIDE, Tripwire)
    - Regular audit of cron job permissions
    - Monitor for unauthorized cron job modifications
    - Use SELinux/AppArmor for additional access controls

**Impact on Users:**

- **Positive**: Prevents privilege escalation, protects system integrity, ensures compliance with security standards
- **Negative**: Requires root access to modify cron jobs, may complicate automated deployment scripts
- **Minimal disruption**: Normal system operation continues, only administrative access is restricted

**Security vs User Friendliness Balance:**

The optimal approach is **strict access control with proper documentation**:

- **Primary action**: Implement strict file permissions (600) with root ownership
- **Administrative access**: Provide clear procedures for authorized cron job management
- **Monitoring**: Implement change detection and alerting for critical system files
- **Documentation**: Create guidelines for secure cron job management and troubleshooting

This approach prioritizes system security by preventing unauthorized access to critical system files while maintaining functionality through proper administrative procedures and documentation.

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
