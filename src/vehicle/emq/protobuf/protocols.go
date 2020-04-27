package protobuf

func GetFlowProtocols(typeInt int) string {

	p := map[int]string{
		0:  "UNKNOWN",
		1:  "FTP_CONTROL",
		2:  "MAIL_POP",
		3:  "MAIL_SMTP",
		4:  "MAIL_IMAP",
		5:  "DNS",
		6:  "IPP",
		7:  "HTTP",
		8:  "MDNS",
		9:  "NTP",
		10: "NETBIOS",
		11: "NFS",
		12: "SSDP",
		13: "BGP",
		14: "SNMP",
		15: "XDMCP",
		16: "SMBV1",
		17: "SYSLOG",
		18: "DHCP",
		19: "POSTGRES",
		20: "MYSQL",
		21: "HOTMAIL",
		22: "DIRECT_DOWNLOAD_LINK",
		23: "MAIL_POPS",
		24: "APPLEJUICE",
		25: "DIRECTCONNECT",
		26: "DIRECTCONNECT",
		27: "COAP",
		28: "VMWARE",
		29: "MAIL_SMTPS",
		30: "FBZERO",
		31: "UBNTAC2",
		32: "KONTIKI",
		33: "OPENFT",
		34: "FASTTRACK",
		35: "GNUTELLA",
	}

	return p[typeInt]

}

//EDONKEY               : 36, /* Tomasz Bujlow <tomasz@skatnet.dk> */
//BITTORRENT            : 37,
//SKYPE_CALL            : 38, /* Skype call and videocalls */
//SIGNAL                : 39,
//MEMCACHED             : 40, /* Memcached - Darryl Sokoloski <darryl@egloo.ca> */
//SMBV23                : 41, /* SMB version 2/3 */
//MINING                : 42, /* Bitcoin, Ethereum, ZCash, Monero */
//NEST_LOG_SINK         : 43, /* Nest Log Sink (Nest Protect) - Darryl Sokoloski <darryl@egloo.ca> */
//MODBUS                : 44, /* Modbus */
//WHATSAPP_CALL         : 45, /* WhatsApp video ad audio calls go here */
//DATASAVER             : 46, /* Protocols used to save data on Internet communications */
//XBOX                  : 47,
//QQ                    : 48,
//TIKTOK                : 49,
//RTSP                  : 50,
//MAIL_IMAPS            : 51,
//ICECAST               : 52,
//PPLIVE                : 53, /* Tomasz Bujlow <tomasz@skatnet.dk> */
//PPSTREAM              : 54,
//ZATTOO                : 55,
//SHOUTCAST             : 56,
//SOPCAST               : 57,
//TVANTS                : 58,
//TVUPLAYER             : 59,
//HTTP_DOWNLOAD         : 60,
//QQLIVE                : 61,
//THUNDER               : 62,
//SOULSEEK              : 63,
//PS_VUE                : 64,
//IRC                   : 65,
//AYIYA                 : 66,
//UNENCRYPTED_JABBER    : 67,
//MSN                   : 68,
//OSCAR                 : 69,
//YAHOO                 : 70,
//BATTLEFIELD           : 71,
//GOOGLE_PLUS           : 72,
//IP_VRRP               : 73,
//STEAM                 : 74, /* Tomasz Bujlow <tomasz@skatnet.dk> */
//HALFLIFE2             : 75,
//WORLDOFWARCRAFT       : 76,
//TELNET                : 77,
//STUN                  : 78,
//IP_IPSEC              : 79,
//IP_GRE                : 80,
//IP_ICMP               : 81, // Internet控制报文协议
//IP_IGMP               : 82, // 路由协议
//IP_EGP                : 83,
//IP_SCTP               : 84,
//IP_OSPF               : 85,
//IP_IP_IN_IP           : 86,
//RTP                   : 87,
//RDP                   : 88,
//VNC                   : 89,
//PCANYWHERE            : 90,
//TLS                   : 91,
//SSH                   : 92,
//USENET                : 93,
//MGCP                  : 94,
//IAX                   : 95,
//TFTP                  : 96,
//AFP                   : 97,
//STEALTHNET            : 98,
//AIMINI                : 99,
//SIP                   : 100,
//TRUPHONE              : 101,
//IP_ICMPV6             : 102,
//DHCPV6                : 103,
//ARMAGETRON            : 104,
//CROSSFIRE             : 105,
//DOFUS                 : 106,
//FIESTA                : 107,
//FLORENSIA             : 108,
//GUILDWARS             : 109,
//HTTP_ACTIVESYNC       : 110,
//KERBEROS              : 111,
//LDAP                  : 112,
//MAPLESTORY            : 113,
//MSSQL_TDS             : 114,
//PPTP                  : 115,
//WARCRAFT3             : 116,
//WORLD_OF_KUNG_FU      : 117,
//SLACK                 : 118,
//FACEBOOK              : 119,
//TWITTER               : 120,
//DROPBOX               : 121,
//GMAIL                 : 122,
//GOOGLE_MAPS           : 123,
//YOUTUBE               : 124,
//SKYPE                 : 125,
//GOOGLE                : 126,
//DCERPC                : 127,
//NETFLOW               : 128,
//SFLOW                 : 129,
//HTTP_CONNECT          : 130,
//HTTP_PROXY            : 131,
//CITRIX                : 132, /* It also includes the old  CITRIX_ONLINE */
//NETFLIX               : 133,
//LASTFM                : 134,
//WAZE                  : 135,
//YOUTUBE_UPLOAD        : 136, /* Upload files to youtube */
//HULU                  : 137,
//CHECKMK               : 138,
//AJP                   : 139, /* Leonn Paiva <leonn.paiva@gmail.com> */
//APPLE                 : 140,
//WEBEX                 : 141,
//WHATSAPP              : 142,
//APPLE_ICLOUD          : 143,
//VIBER                 : 144,
//APPLE_ITUNES          : 145,
//RADIUS                : 146,
//WINDOWS_UPDATE        : 147,
//TEAMVIEWER            : 148, /* xplico.org */
//TUENTI                : 149,
//LOTUS_NOTES           : 150,
//SAP                   : 151,
//GTP                   : 152,
//UPNP                  : 153,
//LLMNR                 : 154,
//REMOTE_SCAN           : 155,
//SPOTIFY               : 156,
//MESSENGER             : 157,
//H323                  : 158, /* Remy Mudingay <mudingay@ill.fr> */
//OPENVPN               : 159, /* Remy Mudingay <mudingay@ill.fr> */
//NOE                   : 160, /* Remy Mudingay <mudingay@ill.fr> */
//CISCOVPN              : 161, /* Remy Mudingay <mudingay@ill.fr> */
//TEAMSPEAK             : 162, /* Remy Mudingay <mudingay@ill.fr> */
//TOR                   : 163, /* Remy Mudingay <mudingay@ill.fr> */
//SKINNY                : 164, /* Remy Mudingay <mudingay@ill.fr> */
//RTCP                  : 165, /* Remy Mudingay <mudingay@ill.fr> */
//RSYNC                 : 166, /* Remy Mudingay <mudingay@ill.fr> */
//ORACLE                : 167, /* Remy Mudingay <mudingay@ill.fr> */
//CORBA                 : 168, /* Remy Mudingay <mudingay@ill.fr> */
//UBUNTUONE             : 169, /* Remy Mudingay <mudingay@ill.fr> */
//WHOIS_DAS             : 170,
//COLLECTD              : 171,
//SOCKS                 : 172, /* Tomasz Bujlow <tomasz@skatnet.dk> */
//NINTENDO              : 173,
//RTMP                  : 174, /* Tomasz Bujlow <tomasz@skatnet.dk> */
//FTP_DATA              : 175, /* Tomasz Bujlow <tomasz@skatnet.dk> */
//WIKIPEDIA             : 176, /* Tomasz Bujlow <tomasz@skatnet.dk> */
//ZMQ                   : 177,
//AMAZON                : 178, /* Tomasz Bujlow <tomasz@skatnet.dk> */
//EBAY                  : 179, /* Tomasz Bujlow <tomasz@skatnet.dk> */
//CNN                   : 180, /* Tomasz Bujlow <tomasz@skatnet.dk> */
//MEGACO                : 181, /* Gianluca Costa <g.costa@xplico.org> */
//REDIS                 : 182,
//PANDO                 : 183, /* Tomasz Bujlow <tomasz@skatnet.dk> */
//VHUA                  : 184,
//TELEGRAM              : 185, /* Gianluca Costa <g.costa@xplico.org> */
//VEVO                  : 186,
//PANDORA               : 187,
//QUIC                  : 188, /* Andrea Buscarinu <andrea.buscarinu@gmail.com> - Michele Campus <michelecampus5@gmail.com> */
//ZOOM                  : 189, /* Zoom video conference. */
//EAQ                   : 190,
//OOKLA                 : 191,
//AMQP                  : 192,
//KAKAOTALK             : 193, /* KakaoTalk Chat (no voice call) */
//KAKAOTALK_VOICE       : 194, /* KakaoTalk Voice */
//TWITCH                : 195, /* Edoardo Dominici <edoaramis@gmail.com> */
//DOH_DOT               : 196, /* DoH (DNS over HTTPS), DoT (DNS over TLS) */
//WECHAT                : 197,
//MPEGTS                : 198,
//SNAPCHAT              : 199,
//SINA                  : 200,
//HANGOUT_DUO           : 201, /* Google Hangout ad Duo (merged as they are very similar) */
//IFLIX                 : 202, /* www.vizuamatix.com R&D team & M.Mallawaarachchie <manoj_ws@yahoo.com> */
//GITHUB                : 203,
//BJNP                  : 204,
//FREE_205              : 205,
//WIREGUARD             : 206,
//SMPP                  : 207, /* Damir Franusic <df@release14.org> */
//DNSCRYPT              : 208,
//TINC                  : 209, /* William Guglielmo <william@deselmo.com> */
//DEEZER                : 210,
//INSTAGRAM             : 211, /* Andrea Buscarinu <andrea.buscarinu@gmail.com> */
//MICROSOFT             : 212,
//STARCRAFT             : 213, /* Matteo Bracci <matteobracci1@gmail.com> */
//TEREDO                : 214,
//HOTSPOT_SHIELD        : 215,
//IMO                   : 216,
//GOOGLE_DRIVE          : 217,
//OCS                   : 218,
//OFFICE_365            : 219,
//CLOUDFLARE            : 220,
//MS_ONE_DRIVE          : 221,
//MQTT                  : 222,
//RX                    : 223,
//APPLESTORE            : 224,
//OPENDNS               : 225,
//GIT                   : 226,
//DRDA                  : 227,
//PLAYSTORE             : 228,
//SOMEIP                : 229,
//FIX                   : 230,
//PLAYSTATION           : 231,
//PASTEBIN              : 232, /* Paulo Angelo <pa@pauloangelo.com> */
//LINKEDIN              : 233, /* Paulo Angelo <pa@pauloangelo.com> */
//SOUNDCLOUD            : 234,
//CSGO                  : 235, /* Counter-Strike Global Offensive, Dota : 2 */
//LISP	                : 236,
//DIAMETER	            : 237,
//APPLE_PUSH            : 238,
//GOOGLE_SERVICES       : 239,
//AMAZON_VIDEO          : 240,
//GOOGLE_DOCS           : 241,
//WHATSAPP_FILES        : 242, /* Videos, pictures, voice messages... */
//TARGUS_GETDATA        : 243,
//DNP3                  : 244,
//IEC60870              : 245, /* https://en.wikipedia.org/wiki/IEC_60870-5 */
//BLOOMBERG             : 246,
//CAPWAP                : 247,
//ZABBIX                : 248,
