package protobuf

func GetFlowProtocols(typeInt int) string {
	p := map[int]string{
		0:   "UNKNOWN",
		1:   "FTP_CONTROL",
		2:   "MAIL_POP",
		3:   "MAIL_SMTP",
		4:   "MAIL_IMAP",
		5:   "DNS",
		6:   "IPP",
		7:   "HTTP",
		8:   "MDNS",
		9:   "NTP",
		10:  "NETBIOS",
		11:  "NFS",
		12:  "SSDP",
		13:  "BGP",
		14:  "SNMP",
		15:  "XDMCP",
		16:  "SMBV1",
		17:  "SYSLOG",
		18:  "DHCP",
		19:  "POSTGRES",
		20:  "MYSQL",
		21:  "HOTMAIL",
		22:  "DIRECT_DOWNLOAD_LINK",
		23:  "MAIL_POPS",
		24:  "APPLEJUICE",
		25:  "DIRECTCONNECT",
		26:  "NTOP",
		27:  "COAP",
		28:  "VMWARE",
		29:  "MAIL_SMTPS",
		30:  "FBZERO",
		31:  "UBNTAC2",
		32:  "KONTIKI",
		33:  "OPENFT",
		34:  "FASTTRACK",
		35:  "GNUTELLA",
		36:  "EDONKEY",
		37:  "BITTORRENT",
		38:  "SKYPE_CALL",
		39:  "SIGNAL",
		40:  "MEMCACHED",
		41:  "SMBV23",
		42:  "MINING",
		43:  "NEST_LOG_SINK",
		44:  "MODBUS",
		45:  "WHATSAPP_CALL",
		46:  "DATASAVER",
		47:  "XBOX",
		48:  "QQ",
		49:  "TIKTOK",
		50:  "RTSP",
		51:  "MAIL_IMAPS",
		52:  "ICECAST",
		53:  "PPLIVE",
		54:  "PPSTREAM",
		55:  "ZATTOO",
		56:  "SHOUTCAST",
		57:  "SOPCAST",
		58:  "TVANTS",
		59:  "TVUPLAYER",
		60:  "HTTP_DOWNLOAD",
		61:  "QQLIVE",
		62:  "THUNDER",
		63:  "SOULSEEK",
		64:  "PS_VUE",
		65:  "IRC",
		66:  "AYIYA",
		67:  "UNENCRYPTED_JABBER",
		68:  "MSN",
		69:  "OSCAR",
		70:  "YAHOO",
		71:  "BATTLEFIELD",
		72:  "GOOGLE_PLUS",
		73:  "IP_VRRP",
		74:  "STEAM",
		75:  "HALFLIFE2",
		76:  "WORLDOFWARCRAFT",
		77:  "TELNET",
		78:  "STUN",
		79:  "IP_IPSEC",
		80:  "IP_GRE",
		81:  "IP_ICMP",
		82:  "IP_IGMP",
		83:  "IP_EGP",
		84:  "IP_SCTP",
		85:  "IP_OSPF",
		86:  "IP_IP_IN_IP",
		87:  "RTP",
		88:  "RDP",
		89:  "VNC",
		90:  "PCANYWHERE",
		91:  "TLS",
		92:  "SSH",
		93:  "USENET",
		94:  "MGCP",
		95:  "IAX",
		96:  "TFTP",
		97:  "AFP",
		98:  "STEALTHNET",
		99:  "AIMINI",
		100: "SIP",
		101: "TRUPHONE",
		102: "IP_ICMPV6",
		103: "DHCPV6",
		104: "ARMAGETRON",
		105: "CROSSFIRE",
		106: "DOFUS",
		107: "FIESTA",
		108: "FLORENSIA",
		109: "GUILDWARS",
		110: "HTTP_ACTIVESYNC",
		111: "KERBEROS",
		112: "LDAP",
		113: "MAPLESTORY",
		114: "MSSQL_TDS",
		115: "PPTP",
		116: "WARCRAFT3",
		117: "WORLD_OF_KUNG_FU",
		118: "SLACK",
		119: "FACEBOOK",
		120: "TWITTER",
		121: "DROPBOX",
		122: "GMAIL",
		123: "GOOGLE_MAPS",
		124: "YOUTUBE",
		125: "SKYPE",
		126: "GOOGLE",
		127: "DCERPC",
		128: "NETFLOW",
		129: "SFLOW",
		130: "HTTP_CONNECT",
		131: "HTTP_PROXY",
		132: "CITRIX",
		133: "NETFLIX",
		134: "LASTFM",
		135: "WAZE",
		136: "YOUTUBE_UPLOAD",
		137: "HULU",
		138: "CHECKMK",
		139: "AJP",
		140: "APPLE",
		141: "WEBEX",
		142: "WHATSAPP",
		143: "APPLE_ICLOUD",
		144: "VIBER",
		145: "APPLE_ITUNES",
		146: "RADIUS",
		147: "WINDOWS_UPDATE",
		148: "TEAMVIEWER",
		149: "TUENTI",
		150: "LOTUS_NOTES",
		151: "SAP",
		152: "GTP",
		153: "UPNP",
		154: "LLMNR",
		155: "REMOTE_SCAN",
		156: "SPOTIFY",
		157: "MESSENGER",
		158: "H323",
		159: "OPENVPN",
		160: "NOE",
		161: "CISCOVPN",
		162: "TEAMSPEAK",
		163: "TOR",
		164: "SKINNY",
		165: "RTCP",
		166: "RSYNC",
		167: "ORACLE",
		168: "CORBA",
		169: "UBUNTUONE",
		170: "WHOIS_DAS",
		171: "COLLECTD",
		172: "SOCKS",
		173: "NINTENDO",
		174: "RTMP",
		175: "FTP_DATA",
		176: "WIKIPEDIA",
		177: "ZMQ",
		178: "AMAZON",
		179: "EBAY",
		180: "CNN",
		181: "MEGACO",
		182: "REDIS",
		183: "PANDO",
		184: "VHUA",
		185: "TELEGRAM",
		186: "VEVO",
		187: "PANDORA",
		188: "QUIC",
		189: "ZOOM",
		190: "EAQ",
		191: "OOKLA",
		192: "AMQP",
		193: "KAKAOTALK",
		194: "KAKAOTALK_VOICE",
		195: "TWITCH",
		196: "DOH_DOT",
		197: "WECHAT",
		198: "MPEGTS",
		199: "SNAPCHAT",
		200: "SINA",
		201: "HANGOUT_DUO",
		202: "IFLIX",
		203: "GITHUB",
		204: "BJNP",
		205: "FREE_205",
		206: "WIREGUARD",
		207: "SMPP",
		208: "DNSCRYPT",
		209: "TINC",
		210: "DEEZER",
		211: "INSTAGRAM",
		212: "MICROSOFT",
		213: "STARCRAFT",
		214: "TEREDO",
		215: "HOTSPOT_SHIELD",
		216: "IMO",
		217: "GOOGLE_DRIVE",
		218: "OCS",
		219: "OFFICE_365",
		220: "CLOUDFLARE",
		221: "MS_ONE_DRIVE",
		222: "MQTT",
		223: "RX",
		224: "APPLESTORE",
		225: "OPENDNS",
		226: "GIT",
		227: "DRDA",
		228: "PLAYSTORE",
		229: "SOMEIP",
		230: "FIX",
		231: "PLAYSTATION",
		232: "PASTEBIN",
		233: "LINKEDIN",
		234: "SOUNDCLOUD",
		235: "CSGO",
		236: "LISP",
		237: "DIAMETER",
		238: "APPLE_PUSH",
		239: "GOOGLE_SERVICES",
		240: "AMAZON_VIDEO",
		241: "GOOGLE_DOCS",
		242: "WHATSAPP_FILES",
		243: "TARGUS_GETDATA",
		244: "DNP3",
		245: "IEC60870",
		246: "BLOOMBERG",
		247: "CAPWAP",
		248: "ZABBIX",
	}
	return p[typeInt]
}
