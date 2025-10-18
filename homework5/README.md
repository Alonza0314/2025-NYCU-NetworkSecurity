# Homework 5

> [!Note]
> The system is **Ubuntu 20.04**.

## Analysis Table

- [1 Initial Setup](#1-initial-setup)

    - [1.1.1.1 Ensure cramfs kernel module is not available](#1111-ensure-cramfs-kernel-module-is-not-available)
    - [1.1.1.2 Ensure freevxfs kernel module is not available](#1112-ensure-freevxfs-kernel-module-is-not-available)
    - [1.1.1.3 Ensure hfs kernel module is not available](#1113-ensure-hfs-kernel-module-is-not-available)
    - [1.1.1.4 Ensure hfsplus kernel module is not available](#1114-ensure-hfsplus-kernel-module-is-not-available)

- [2 Services](#2-services)

    - [2.1.13 Ensure rsync services are not in use](#2113-ensure-rsync-services-are-not-in-use)
    - [2.2.4 Ensure telnet client is not installed](#224-ensure-telnet-client-is-not-installed)
    - [2.2.6 Ensure ftp client is not installed](#226-ensure-ftp-client-is-not-installed)
    - [2.3.2.1 Ensure systemd-timesyncd configured with authorized timeserver](#2321-ensure-systemd-timesyncd-configured-with-authorized-timeserver)

- [3 Network](#3-network)

    - [3.3.1 Ensure ip forwarding is disabled](#331-ensure-ip-forwarding-is-disabled)
    - [3.3.2 Ensure packet redirect sending is disabled](#332-ensure-packet-redirect-sending-is-disabled)
    - [3.3.5 Ensure icmp redirects are not accepted](#335-ensure-icmp-redirects-are-not-accepted)
    - [3.3.11 Ensure ipv6 router advertisements are not accepted](#3311-ensure-ipv6-router-advertisements-are-not-accepted)

- [4 Host Based Firewall](#4-host-based-firewall)
- [5 Access Control](#5-access-control)

## 1 Initial Setup

### 1.1.1.1 Ensure cramfs kernel module is not available

**Analysis:**

**Why OS left default configurations:**

- **Legacy Compatibility**: cramfs is an older, compressed read-only filesystem, historically used in embedded systems or boot images where space was extremely limited
- **Specific Use Cases**: Might have been included for niche applications requiring a highly compressed, immutable filesystem
- **Minimal Overhead**: Its read-only nature and compression offered benefits for specific, resource-constrained environments
- **Ease-of-Use**: Pre-installed kernel module support for various filesystem types without manual configuration

**Security Issues:**

- **Unnecessary Attack Surface**: If cramfs is not actively used, having its kernel module available increases the system's attack surface
- **Potential Vulnerabilities**: As an older and less actively maintained filesystem, cramfs might contain undiscovered vulnerabilities that could be exploited
- **Kernel Exploitation**: A vulnerability in the cramfs module could potentially lead to kernel-level exploits, compromising the entire system
- **Compliance violations**: Many security standards require disabling unused kernel modules to reduce attack surface

**Remediation Solutions:**

1. **Unload the cramfs kernel module**:

    ```bash
    # Unload cramfs module if currently loaded
    modprobe -r cramfs 2>/dev/null
    rmmod cramfs 2>/dev/null
    ```

2. **Disable cramfs kernel module permanently**:

    ```bash
    # Prevent cramfs from being loaded by making any attempt to install it fail
    printf '\n%s\n' "install cramfs /bin/false" >> /etc/modprobe.d/cramfs.conf
    
    # Blacklist cramfs to prevent automatic loading
    printf '\n%s\n' "blacklist cramfs" >> /etc/modprobe.d/cramfs.conf
    
    # Update initramfs to ensure changes persist across reboots
    update-initramfs -u
    ```

3. **Verify configuration**:

    ```bash
    # Check if cramfs module is loaded
    lsmod | grep cramfs
    
    # Verify modprobe configuration
    cat /etc/modprobe.d/cramfs.conf
    ```

**Impact on Users:**

- **Positive**: Reduces attack surface, improves system security, prevents potential kernel exploits
- **Negative**: Disables cramfs functionality (rarely used in modern systems), may break specific embedded applications
- **Minimal disruption**: For typical desktop/workstation systems, no impact on normal operations

**Security vs User Friendliness Balance:**

The optimal approach is **disable unless specifically required**:

- **Primary action**: Disable cramfs kernel module on general-purpose systems
- **Exception handling**: Enable only when cramfs is explicitly needed for specific applications
- **Documentation**: Clearly document when and why cramfs should be enabled
- **Monitoring**: Implement alerts for unauthorized kernel module loading

This approach prioritizes security by reducing the attack surface while maintaining functionality for systems that legitimately require cramfs support.

### 1.1.1.2 Ensure freevxfs kernel module is not available

**Analysis:**

**Why OS left default configurations:**

- **Legacy Compatibility**: freevxfs is a free version of the Veritas filesystem, historically used for HP-UX operating systems compatibility
- **Cross-platform Support**: Included to support filesystem interoperability between different Unix-like systems
- **Specific Use Cases**: Might have been included for enterprise environments requiring HP-UX filesystem access
- **Ease-of-Use**: Pre-installed kernel module support for various filesystem types without manual configuration

**Security Issues:**

- **Unnecessary Attack Surface**: If freevxfs is not actively used, having its kernel module available increases the system's attack surface
- **Potential Vulnerabilities**: As a less commonly used filesystem, freevxfs might contain undiscovered vulnerabilities that could be exploited
- **Kernel Exploitation**: A vulnerability in the freevxfs module could potentially lead to kernel-level exploits, compromising the entire system
- **Compliance violations**: Many security standards require disabling unused kernel modules to reduce attack surface

**Remediation Solutions:**

1. **Unload the freevxfs kernel module**:

    ```bash
    # Unload freevxfs module if currently loaded
    modprobe -r freevxfs 2>/dev/null
    rmmod freevxfs 2>/dev/null
    ```

2. **Disable freevxfs kernel module permanently**:

    ```bash
    # Prevent freevxfs from being loaded by making any attempt to install it fail
    printf '\n%s\n' "install freevxfs /bin/false" >> /etc/modprobe.d/freevxfs.conf
    
    # Blacklist freevxfs to prevent automatic loading
    printf '\n%s\n' "blacklist freevxfs" >> /etc/modprobe.d/freevxfs.conf
    
    # Update initramfs to ensure changes persist across reboots
    update-initramfs -u
    ```

3. **Verify configuration**:

    ```bash
    # Check if freevxfs module is loaded
    lsmod | grep freevxfs
    
    # Verify modprobe configuration
    cat /etc/modprobe.d/freevxfs.conf
    ```

**Impact on Users:**

- **Positive**: Reduces attack surface, improves system security, prevents potential kernel exploits
- **Negative**: Disables freevxfs functionality (rarely used in modern Linux systems), may break HP-UX filesystem access
- **Minimal disruption**: For typical desktop/workstation systems, no impact on normal operations

**Security vs User Friendliness Balance:**

The optimal approach is **disable unless specifically required**:

- **Primary action**: Disable freevxfs kernel module on general-purpose systems
- **Exception handling**: Enable only when freevxfs is explicitly needed for HP-UX compatibility
- **Documentation**: Clearly document when and why freevxfs should be enabled
- **Monitoring**: Implement alerts for unauthorized kernel module loading

This approach prioritizes security by reducing the attack surface while maintaining functionality for systems that legitimately require freevxfs support for HP-UX filesystem access.

### 1.1.1.3 Ensure hfs kernel module is not available

**Analysis:**

**Why OS left default configurations:**

- **Cross-platform Compatibility**: HFS (Hierarchical File System) is used to mount Mac OS filesystems, enabling interoperability between Linux and macOS systems
- **Data Migration Support**: Included to support users migrating data from Mac systems or accessing Mac-formatted storage devices
- **Specific Use Cases**: Might have been included for enterprise environments requiring Mac filesystem access
- **Ease-of-Use**: Pre-installed kernel module support for various filesystem types without manual configuration

**Security Issues:**

- **Unnecessary Attack Surface**: If HFS is not actively used, having its kernel module available increases the system's attack surface
- **Potential Vulnerabilities**: As a less commonly used filesystem in Linux environments, HFS might contain undiscovered vulnerabilities that could be exploited
- **Kernel Exploitation**: A vulnerability in the HFS module could potentially lead to kernel-level exploits, compromising the entire system
- **Compliance violations**: Many security standards require disabling unused kernel modules to reduce attack surface

**Remediation Solutions:**

1. **Unload the HFS kernel module**:

    ```bash
    # Unload HFS module if currently loaded
    modprobe -r hfs 2>/dev/null
    rmmod hfs 2>/dev/null
    ```

2. **Disable HFS kernel module permanently**:

    ```bash
    # Prevent HFS from being loaded by making any attempt to install it fail
    printf '\n%s\n' "install hfs /bin/false" >> /etc/modprobe.d/hfs.conf
    
    # Blacklist HFS to prevent automatic loading
    printf '\n%s\n' "blacklist hfs" >> /etc/modprobe.d/hfs.conf
    
    # Update initramfs to ensure changes persist across reboots
    update-initramfs -u
    ```

3. **Verify configuration**:

    ```bash
    # Check if HFS module is loaded
    lsmod | grep hfs
    
    # Verify modprobe configuration
    cat /etc/modprobe.d/hfs.conf
    ```

**Impact on Users:**

- **Positive**: Reduces attack surface, improves system security, prevents potential kernel exploits
- **Negative**: Disables HFS functionality, may break Mac filesystem access and data migration capabilities
- **Minimal disruption**: For typical desktop/workstation systems without Mac interaction, no impact on normal operations

**Security vs User Friendliness Balance:**

The optimal approach is **disable unless specifically required**:

- **Primary action**: Disable HFS kernel module on general-purpose systems
- **Exception handling**: Enable only when HFS is explicitly needed for Mac filesystem access
- **Documentation**: Clearly document when and why HFS should be enabled
- **Monitoring**: Implement alerts for unauthorized kernel module loading

This approach prioritizes security by reducing the attack surface while maintaining functionality for systems that legitimately require HFS support for Mac filesystem access.

### 1.1.1.4 Ensure hfsplus kernel module is not available

**Analysis:**

**Why OS left default configurations:**

- **Cross-platform Compatibility**: HFS+ (Hierarchical File System Plus) is the successor to HFS, designed to replace HFS and enable mounting of modern Mac OS filesystems
- **Data Migration Support**: Included to support users migrating data from modern Mac systems or accessing Mac-formatted storage devices
- **Specific Use Cases**: Might have been included for enterprise environments requiring modern Mac filesystem access
- **Ease-of-Use**: Pre-installed kernel module support for various filesystem types without manual configuration

**Security Issues:**

- **Unnecessary Attack Surface**: If HFS+ is not actively used, having its kernel module available increases the system's attack surface
- **Potential Vulnerabilities**: As a less commonly used filesystem in Linux environments, HFS+ might contain undiscovered vulnerabilities that could be exploited
- **Kernel Exploitation**: A vulnerability in the HFS+ module could potentially lead to kernel-level exploits, compromising the entire system
- **Compliance violations**: Many security standards require disabling unused kernel modules to reduce attack surface

**Remediation Solutions:**

1. **Unload the HFS+ kernel module**:

    ```bash
    # Unload HFS+ module if currently loaded
    modprobe -r hfsplus 2>/dev/null
    rmmod hfsplus 2>/dev/null
    ```

2. **Disable HFS+ kernel module permanently**:

    ```bash
    # Prevent HFS+ from being loaded by making any attempt to install it fail
    printf '\n%s\n' "install hfsplus /bin/false" >> /etc/modprobe.d/hfsplus.conf
    
    # Blacklist HFS+ to prevent automatic loading
    printf '\n%s\n' "blacklist hfsplus" >> /etc/modprobe.d/hfsplus.conf
    
    # Update initramfs to ensure changes persist across reboots
    update-initramfs -u
    ```

3. **Verify configuration**:

    ```bash
    # Check if HFS+ module is loaded
    lsmod | grep hfsplus
    
    # Verify modprobe configuration
    cat /etc/modprobe.d/hfsplus.conf
    ```

**Impact on Users:**

- **Positive**: Reduces attack surface, improves system security, prevents potential kernel exploits
- **Negative**: Disables HFS+ functionality, may break modern Mac filesystem access and data migration capabilities
- **Minimal disruption**: For typical desktop/workstation systems without Mac interaction, no impact on normal operations

**Security vs User Friendliness Balance:**

The optimal approach is **disable unless specifically required**:

- **Primary action**: Disable HFS+ kernel module on general-purpose systems
- **Exception handling**: Enable only when HFS+ is explicitly needed for modern Mac filesystem access
- **Documentation**: Clearly document when and why HFS+ should be enabled
- **Monitoring**: Implement alerts for unauthorized kernel module loading

This approach prioritizes security by reducing the attack surface while maintaining functionality for systems that legitimately require HFS+ support for modern Mac filesystem access.

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

## 3 Network

### 3.3.1 Ensure ip forwarding is disabled

**Analysis:**

**Why OS left default configurations:**

- **Service Reachability**: IP forwarding enables routing functionality, allowing the system to act as a router or gateway
- **User Friendliness**: Default enabled state supports network connectivity and multi-homed systems out-of-the-box
- **Ease-of-Use**: Automatic routing capabilities reduce manual network configuration complexity
- **Functionality**: Enables advanced networking features like NAT, VPN, and network bridging

**Security Issues:**

- **Attack surface expansion**: Enables the system to route traffic, increasing potential attack vectors
- **Network reconnaissance**: Allows attackers to use the system as a pivot point for network scanning
- **Traffic interception**: System can intercept and potentially modify routed network traffic
- **Lateral movement**: Facilitates attackers moving between network segments
- **Compliance violations**: Many security standards require disabling IP forwarding on non-router systems

**Remediation Solutions:**

1. **Disable IP forwarding**:

    ```bash
    # Check current status
    sysctl net.ipv4.ip_forward
    
    # Disable IP forwarding
    echo 'net.ipv4.ip_forward = 0' >> /etc/sysctl.conf
    
    # Apply immediately
    sysctl -p
    
    # Verify configuration
    sysctl net.ipv4.ip_forward
    ```

2. **Disable IPv6 forwarding** (if applicable):

    ```bash
    # Disable IPv6 forwarding
    echo 'net.ipv6.conf.all.forwarding = 0' >> /etc/sysctl.conf
    sysctl -p
    ```

3. **Verify configuration**:

    ```bash
    # Check all forwarding settings
    sysctl -a | grep forwarding
    
    # Ensure both IPv4 and IPv6 forwarding are disabled
    sysctl net.ipv4.ip_forward net.ipv6.conf.all.forwarding
    ```

**Impact on Users:**

- **Positive**: Reduces attack surface, prevents unauthorized routing, improves security posture
- **Negative**: Disables routing functionality, may break VPN or NAT configurations, affects multi-homed systems
- **Minimal disruption**: For typical desktop/workstation systems, no impact on normal operations

**Security vs User Friendliness Balance:**

The optimal approach is **disable unless specifically required**:

- **Primary action**: Disable IP forwarding on non-router systems
- **Exception handling**: Enable only when routing functionality is explicitly needed
- **Documentation**: Clearly document when and why IP forwarding should be enabled
- **Monitoring**: Implement alerts for unauthorized IP forwarding activation

This approach prioritizes security by default while maintaining functionality for systems that legitimately require routing capabilities.

### 3.3.2 Ensure packet redirect sending is disabled

**Analysis:**

**Why OS left default configurations:**

- **Service Reachability**: ICMP redirects help optimize network routing by informing hosts of better routes
- **User Friendliness**: Default enabled state supports automatic network optimization without manual configuration
- **Ease-of-Use**: Automatic route optimization reduces network configuration complexity
- **Legacy compatibility**: Historical network implementations relied on ICMP redirects for optimal routing

**Security Issues:**

- **Attack surface expansion**: Enables malicious ICMP redirect attacks that can disrupt network routing
- **Route manipulation**: Attackers can use compromised hosts to send malicious ICMP redirects
- **Network disruption**: Malicious redirects can cause network routing failures and connectivity issues
- **Compliance violations**: Many security standards require disabling ICMP redirects on non-router systems
- **Lateral movement**: Facilitates attackers manipulating network traffic flow

**Remediation Solutions:**

1. **Disable ICMP redirect sending**:

    ```bash
    # Create configuration file
    printf '%s\n' "net.ipv4.conf.all.send_redirects = 0" "net.ipv4.conf.default.send_redirects = 0" >> /etc/sysctl.d/60-netipv4_sysctl.conf
    
    # Apply configuration immediately
    sysctl --system
    
    # Alternative: Apply specific parameters
    sysctl -w net.ipv4.conf.all.send_redirects=0
    sysctl -w net.ipv4.conf.default.send_redirects=0
    sysctl -w net.ipv4.route.flush=1
    ```

2. **Verify configuration**:

    ```bash
    # Check all interfaces parameter
    sysctl net.ipv4.conf.all.send_redirects
    
    # Check default interface parameter
    sysctl net.ipv4.conf.default.send_redirects
    
    # Expected output: both should show = 0
    ```

3. **Alternative configuration methods**:

    - Add to `/etc/sysctl.conf` if preferred
    - Use higher numbered files in `/etc/sysctl.d/` to override existing settings
    - Implement monitoring to detect unauthorized changes

**Impact on Users:**

- **Positive**: Prevents ICMP redirect attacks, improves network security, reduces attack surface
- **Negative**: Disables automatic route optimization, may require manual network configuration
- **Minimal disruption**: For typical desktop/workstation systems, no impact on normal operations

**Security vs User Friendliness Balance:**

The optimal approach is **disable unless specifically required**:

- **Primary action**: Disable ICMP redirect sending on non-router systems
- **Exception handling**: Enable only when routing optimization is explicitly needed
- **Documentation**: Clearly document when and why ICMP redirects should be enabled
- **Monitoring**: Implement alerts for unauthorized ICMP redirect configuration changes

This approach prioritizes security by preventing malicious ICMP redirect attacks while maintaining functionality for systems that legitimately require routing optimization capabilities.

### 3.3.5 Ensure icmp redirects are not accepted

**Analysis:**

**Why OS left default configurations:**

- **Service Reachability**: ICMP redirects help optimize network routing by allowing external routing devices to update system routing tables
- **User Friendliness**: Default enabled state supports automatic network optimization without manual configuration
- **Ease-of-Use**: Automatic route updates reduce network configuration complexity
- **Legacy compatibility**: Historical network implementations relied on ICMP redirects for optimal routing

**Security Issues:**

- **Man-in-the-middle attacks**: Attackers can send malicious ICMP redirects to manipulate routing tables
- **Route manipulation**: Malicious redirects can redirect traffic through attacker-controlled systems
- **Network disruption**: Malicious redirects can cause network routing failures and connectivity issues
- **Traffic interception**: Attackers can intercept and potentially modify redirected network traffic
- **Compliance violations**: Many security standards require disabling ICMP redirect acceptance on non-router systems

**Remediation Solutions:**

1. **Disable ICMP redirect acceptance**:

    ```bash
    # Disable IPv4 ICMP redirect acceptance
    printf '%s\n' "net.ipv4.conf.all.accept_redirects = 0" "net.ipv4.conf.default.accept_redirects = 0" >> /etc/sysctl.d/60-netipv4_sysctl.conf
    
    # Disable IPv6 ICMP redirect acceptance
    printf '%s\n' "net.ipv6.conf.all.accept_redirects = 0" "net.ipv6.conf.default.accept_redirects = 0" >> /etc/sysctl.d/60-netipv6_sysctl.conf
    
    # Apply configuration immediately
    sysctl --system
    ```

2. **Apply active parameters**:

    ```bash
    # Disable IPv4 ICMP redirect acceptance
    sysctl -w net.ipv4.conf.all.accept_redirects=0
    sysctl -w net.ipv4.conf.default.accept_redirects=0
    sysctl -w net.ipv4.route.flush=1
    
    # Disable IPv6 ICMP redirect acceptance (if IPv6 is enabled)
    sysctl -w net.ipv6.conf.all.accept_redirects=0
    sysctl -w net.ipv6.conf.default.accept_redirects=0
    sysctl -w net.ipv6.route.flush=1
    ```

3. **Verify configuration**:

    ```bash
    # Check IPv4 parameters
    sysctl net.ipv4.conf.all.accept_redirects
    sysctl net.ipv4.conf.default.accept_redirects
    
    # Check IPv6 parameters
    sysctl net.ipv6.conf.all.accept_redirects
    sysctl net.ipv6.conf.default.accept_redirects
    
    # Expected output: all should show = 0
    ```

**Impact on Users:**

- **Positive**: Prevents man-in-the-middle attacks, improves network security, reduces attack surface
- **Negative**: Disables automatic route optimization, may require manual network configuration
- **Minimal disruption**: For typical desktop/workstation systems, no impact on normal operations

**Security vs User Friendliness Balance:**

The optimal approach is **disable unless specifically required**:

- **Primary action**: Disable ICMP redirect acceptance on non-router systems
- **Exception handling**: Enable only when routing optimization is explicitly needed
- **Documentation**: Clearly document when and why ICMP redirect acceptance should be enabled
- **Monitoring**: Implement alerts for unauthorized ICMP redirect configuration changes

This approach prioritizes security by preventing malicious ICMP redirect attacks while maintaining functionality for systems that legitimately require routing optimization capabilities.

### 3.3.11 Ensure ipv6 router advertisements are not accepted

**Analysis:**

**Why OS left default configurations:**

- **Service Reachability**: IPv6 Router Advertisements (RA) enable automatic network configuration and router discovery
- **User Friendliness**: Default enabled state supports automatic IPv6 network configuration without manual setup
- **Ease-of-Use**: Automatic network configuration reduces IPv6 deployment complexity
- **Legacy compatibility**: Historical IPv6 implementations relied on RA for network auto-configuration

**Security Issues:**

- **Router advertisement spoofing**: Attackers can send malicious RA messages to manipulate network configuration
- **Traffic redirection**: Malicious RA messages can redirect traffic through attacker-controlled routers
- **Network disruption**: Malicious RA messages can cause network configuration failures and connectivity issues
- **Man-in-the-middle attacks**: Attackers can use RA spoofing to intercept and modify network traffic
- **Compliance violations**: Many security standards require disabling RA acceptance on non-router systems

**Remediation Solutions:**

1. **Disable IPv6 Router Advertisement acceptance**:

    ```bash
    # Disable IPv6 Router Advertisement acceptance
    printf '%s\n' "net.ipv6.conf.all.accept_ra = 0" "net.ipv6.conf.default.accept_ra = 0" >> /etc/sysctl.d/60-netipv6_sysctl.conf
    
    # Apply configuration immediately
    sysctl --system
    ```

2. **Apply active parameters**:

    ```bash
    # Disable all interfaces from accepting IPv6 Router Advertisements
    sysctl -w net.ipv6.conf.all.accept_ra=0
    
    # Disable default interface from accepting IPv6 Router Advertisements
    sysctl -w net.ipv6.conf.default.accept_ra=0
    
    # Clear IPv6 route cache
    sysctl -w net.ipv6.route.flush=1
    ```

3. **Verify configuration**:

    ```bash
    # Check all interfaces parameter
    sysctl net.ipv6.conf.all.accept_ra
    
    # Check default interface parameter
    sysctl net.ipv6.conf.default.accept_ra
    
    # Expected output: both should show = 0
    ```

**Impact on Users:**

- **Positive**: Prevents router advertisement spoofing, improves IPv6 security, reduces attack surface
- **Negative**: Disables automatic IPv6 network configuration, may require manual IPv6 setup
- **Minimal disruption**: For typical desktop/workstation systems, no impact on normal operations

**Security vs User Friendliness Balance:**

The optimal approach is **disable unless specifically required**:

- **Primary action**: Disable IPv6 Router Advertisement acceptance on non-router systems
- **Exception handling**: Enable only when automatic IPv6 configuration is explicitly needed
- **Documentation**: Clearly document when and why IPv6 RA acceptance should be enabled
- **Monitoring**: Implement alerts for unauthorized IPv6 RA configuration changes

This approach prioritizes security by preventing malicious IPv6 Router Advertisement attacks while maintaining functionality for systems that legitimately require automatic IPv6 network configuration.

## 4 Host Based Firewall

## 5 Access Control

## Contribution Table

| Student ID | Works | Percentage |
| - | - | - |
| 314581015 | Chapter 1 | 20% |
| 313581047 | Chapter 2 | 20% |
| 313581038 | Chapter 3 | 20% |
| 313581055 | Chapter 4 | 20% |
| 412581005 | Chapter 5 | 20% |
