FROM acoshift/go-scratch

ADD entrypoint /entrypoint
COPY index.html /dist/index.html
COPY css /dist/css
COPY js /dist/js
COPY picture /dist/picture
EXPOSE 8080

ENTRYPOINT ["/entrypoint"]
