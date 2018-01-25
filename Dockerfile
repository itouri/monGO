FROM centos:7

# Jenkinsインストールに必要なパッケージインストール
RUN yum -y update
RUN yum -y install wget
RUN yum -y install java-1.8.0-openjdk-devel
RUN yum -y install initscripts


# jenkins yumリポジトリ登録
RUN wget -O /etc/yum.repos.d/jenkins.repo http://pkg.jenkins-ci.org/redhat/jenkins.repo
RUN rpm --import http://pkg.jenkins-ci.org/redhat/jenkins-ci.org.key

# Jenkinsとgitをインストール
RUN yum -y install jenkins
RUN yum -y install git
RUN yum clean all