FROM alpine

ADD k8s-usb-plugin /usr/bin/k8s-usb-plugin

CMD ["k8s-usb-plugin"]
