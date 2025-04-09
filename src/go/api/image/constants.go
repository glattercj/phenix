package image

const POSTBUILD_APT_CLEANUP = `
# --------------------------------------------------- Cleanup ----------------------------------------------------
apt clean || apt-get clean || echo "unable to clean apt cache"
`

const POSTBUILD_NO_ROOT_PASSWD = `
# ---------------------------------------------- No Root Password ------------------------------------------------
sed -i 's/nullok_secure/nullok/' /etc/pam.d/common-auth
sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config
sed -i 's/#PermitEmptyPasswords no/PermitEmptyPasswords yes/' /etc/ssh/sshd_config
sed -i 's/PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config
sed -i 's/PermitEmptyPasswords no/PermitEmptyPasswords yes/' /etc/ssh/sshd_config
passwd -d root
`

const POSTBUILD_PHENIX_HOSTNAME = `
# -------------------------------------------------- Hostname ----------------------------------------------------
echo "phenix" > /etc/hostname
sed -i 's/127.0.1.1 .*/127.0.1.1 phenix/' /etc/hosts
cat > /etc/motd <<EOF

██████╗ ██╗  ██╗███████╗███╗  ██╗██╗██╗  ██╗
██╔══██╗██║  ██║██╔════╝████╗ ██║██║╚██╗██╔╝
██████╔╝███████║█████╗  ██╔██╗██║██║ ╚███╔╝
██╔═══╝ ██╔══██║██╔══╝  ██║╚████║██║ ██╔██╗
██║     ██║  ██║███████╗██║ ╚███║██║██╔╝╚██╗
╚═╝     ╚═╝  ╚═╝╚══════╝╚═╝  ╚══╝╚═╝╚═╝  ╚═╝

EOF
echo "\nBuilt with phenix image on $(date)\n\n" >> /etc/motd
`

const POSTBUILD_PHENIX_BASE = `
# ----------------------------------------------------- Base -----------------------------------------------------
cat > /etc/systemd/system/miniccc.service <<EOF
[Unit]
Description=miniccc
[Service]
ExecStart=/opt/minimega/bin/miniccc -v=false -serial /dev/virtio-ports/cc -logfile /var/log/miniccc.log
[Install]
WantedBy=multi-user.target
EOF
cat > /etc/systemd/system/phenix.service <<EOF
[Unit]
Description=phenix startup service
After=network.target systemd-hostnamed.service
[Service]
Environment=LD_LIBRARY_PATH=/usr/local/lib
ExecStart=/usr/local/bin/phenix-start.sh
RemainAfterExit=true
StandardOutput=journal
Type=oneshot
[Install]
WantedBy=multi-user.target
EOF
mkdir -p /etc/systemd/system/multi-user.target.wants
ln -s /etc/systemd/system/miniccc.service /etc/systemd/system/multi-user.target.wants/miniccc.service
ln -s /etc/systemd/system/phenix.service /etc/systemd/system/multi-user.target.wants/phenix.service
mkdir -p /usr/local/bin
cat > /usr/local/bin/phenix-start.sh <<EOF
#!/bin/bash
for file in /etc/phenix/startup/*; do
	echo \$file
	bash \$file
done
EOF
chmod +x /usr/local/bin/phenix-start.sh
mkdir -p /etc/phenix/startup
`

const POSTBUILD_GUI = `
# ----------------------------------------------------- GUI ------------------------------------------------------
apt-get purge -y gdm3 # messes with no-root-password login
mkdir -p /root/.config/xfce4/
echo "TerminalEmulator=gnome-terminal" > /root/.config/xfce4/helpers.rc
`

const POSTBUILD_PROTONUKE = `
# -------------------------------------------------- Protonuke ---------------------------------------------------
cat > /etc/systemd/system/protonuke.service <<EOF
[Unit]
Description=protonuke
After=network-online.target
Wants=network-online.target
[Service]
EnvironmentFile=/etc/default/protonuke
ExecStart=/opt/minimega/bin/protonuke \$PROTONUKE_ARGS
[Install]
WantedBy=multi-user.target
EOF
mkdir -p /etc/systemd/system/multi-user.target.wants
ln -s /etc/systemd/system/protonuke.service /etc/systemd/system/multi-user.target.wants/protonuke.service
`

const POSTBUILD_ENABLE_DHCP = `
# ----------------------------------------------------- DHCP -----------------------------------------------------
echo "#!/bin/bash\ndhclient" > /etc/init.d/dhcp.sh
chmod +x /etc/init.d/dhcp.sh
update-rc.d dhcp.sh defaults 100
`

var PACKAGES_DEFAULT = []string{
	"curl",
	"ethtool",
	"ncat",
	"net-tools",
	"openssh-server",
	"rsync",
	"ssh",
	"tcpdump",
	"tmux",
	"vim",
	"wget",
}

var PACKAGES_KALI = []string{
	"linux-image-amd64",
	"linux-headers-amd64",
	"default-jdk",
}

var PACKAGES_UBUNTU = []string{
	"linux-image-generic",
	"linux-headers-generic",
}

var PACKAGES_MINGUI = []string{
	"wmctrl",
	"xdotool",
	"xubuntu-desktop",
}

var PACKAGES_MINGUI_KALI = []string{
	"kali-desktop-xfce",
	"wmctrl",
	"xdotool",
}

var PACKAGES_COMPONENTS = []string{
	"main",
	"restricted",
	"universe",
	"multiverse",
}

var PACKAGES_COMPONENTS_KALI = []string{
	"main",
	"contrib",
	"non-free",
	"non-free-firmware",
}
