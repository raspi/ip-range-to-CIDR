# ip-range-to-CIDR

![GitHub All Releases](https://img.shields.io/github/downloads/raspi/ip-range-to-CIDR/total?style=for-the-badge)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/raspi/ip-range-to-CIDR?style=for-the-badge)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/raspi/ip-range-to-CIDR?style=for-the-badge)


Convert IP address range to CIDR.

## Examples

    % ip-range-to-cidr 192.168.0.0 192.168.0.255
    192.168.0.0/24

Figuring out github.com's AS range:

    % systemd-resolve github.com                                                                                          
    github.com: 140.82.118.3                       -- link: phy0

    -- Information acquired via protocol DNS in 45.6ms.
    -- Data is authenticated: no

     % openrdap-client --json --type ip 140.82.118.3 | jq -r '"\(.startAddress) \(.endAddress)"' | xargs ./ip-range-to-cidr
     140.82.112.0/20
