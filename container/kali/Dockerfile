FROM kalilinux/kali-rolling

# These have interactive prompts at install time
# let's specify answers now to let docker run automatically:
RUN echo "kismet-capture-common kismet-capture-common/install-setuid boolean false" | debconf-set-selections
RUN echo "kismet-capture-common kismet-capture-common/install-users string root" | debconf-''set-selections
RUN echo "macchanger macchanger/automatically_run boolean false" | debconf-set-selections
RUN echo "wireshark-common wireshark-common/install-setuid boolean false" | debconf-set-selections
RUN echo "console-setup console-setup/charmap47 select UTF-8" | debconf-set-selections
RUN echo "console-setup console-setup/codeset47 select '. Combined - Latin; Slavic Cyrillic; Hebrew; basic Arabic'" | debconf-set-selections
RUN echo "sslh sslh/inetd_or_standalone select 'from inetd'" | debconf-set-selections

RUN apt-get update && apt install -y kali-linux-headless

# IF YOU WANT IT ALL!
# RUN apt-get -y install kali-linux-large

WORKDIR /CSOL-510

# Tools from the Seed Labs assignments

RUN apt-get install -y bsdextrautils ent

RUN git clone https://github.com/seed-labs/seed-labs.git

CMD ["bash"]
