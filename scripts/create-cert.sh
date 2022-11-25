#!/usr/bin/env bash

openssl \
    req \
    -x509 \
    -nodes \
    -days 365 \
    -newkey rsa:2048 \
    -keyout cert.key \
    -out cert.crt

# Example:
#
# Country Name (2 letter code) [AU]: US
# State or Province Name (full name) [Some-State]: CALIFORNIA.
# Locality Name (eg, city) []: LOS ANGELES.
# Organization Name (eg, company) [Internet Widgits Pty Ltd]: SAMPLE COMPANY
# Organizational Unit Name (eg, section) []: SECURITY DEPARTMENT
# Common Name (e.g. server FQDN or YOUR name) []: SAMPLE_DOMAIN or PUBLIC_IP_ADDRESS
# Email Address []: Enter your email address e.g. info@example.com
