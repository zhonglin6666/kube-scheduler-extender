FROM centos:7.4.1708

RUN cp -a /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && yum install -y net-tools telnet \
    && yum clean all

ADD bin/kube-scheduler-extender /kube-scheduler-extender

ENTRYPOINT ["/kube-scheduler-extender"]