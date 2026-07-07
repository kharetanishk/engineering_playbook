lan - local area netwrok
devices that are connected within an area and communicate with each other
, as each device has their own ip
ip(internet protocol) can be range from 0.0.0.0 to 255.255.255.255 (32bit value) 1bit = 1  or 0
switch - helps in communication between devices within a lan 

now suppose if we want to communicate with the internet (wan) 

then it can happen because of router

Router - it sits between lan and wan , and help connect devices on lan and wan

if you req to facebook
the req travel like

phone  - > router - > facebook server

ip address of the router is gateway

📘 Subnet

A subnet (subnetwork) is a smaller network created by dividing a larger IP network.

Devices in the same subnet can communicate directly.
Devices in different subnets require a router to communicate.
The subnet mask or CIDR prefix determines which part of an IP address is the Network ID and which part is the Host ID.

Example:

IP Address : 192.168.1.25
Subnet Mask: 255.255.255.0 (/24)
Network ID: 192.168.1.0
Host ID: 25
📘 CIDR (Classless Inter-Domain Routing)

CIDR is a shorthand notation that specifies how many bits of an IP address belong to the network.

Format:

IP_Address/Prefix_Length

Examples:

192.168.1.25/24
10.0.0.5/8
172.16.5.20/16
Common CIDR Prefixes
CIDR	Subnet Mask	Network Part	Host Part
/8	255.0.0.0	First 8 bits (1st octet)	Last 24 bits
/16	255.255.0.0	First 16 bits (2 octets)	Last 16 bits
/24	255.255.255.0	First 24 bits (3 octets)	Last 8 bits
🧠 Memory Trick
Subnet = A logical group of devices in the same network.
CIDR (/24, /16, /8) = Tells how many bits are reserved for the Network ID.
Remaining bits = Used for the Host ID (individual devices).