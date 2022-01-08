# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_gandi.config as __config
    config = __config
    import pulumi_gandi.domain as __domain
    domain = __domain
    import pulumi_gandi.email as __email
    email = __email
    import pulumi_gandi.livedns as __livedns
    livedns = __livedns
    import pulumi_gandi.simplehosting as __simplehosting
    simplehosting = __simplehosting
else:
    config = _utilities.lazy_import('pulumi_gandi.config')
    domain = _utilities.lazy_import('pulumi_gandi.domain')
    email = _utilities.lazy_import('pulumi_gandi.email')
    livedns = _utilities.lazy_import('pulumi_gandi.livedns')
    simplehosting = _utilities.lazy_import('pulumi_gandi.simplehosting')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "gandi",
  "mod": "domain/dnssecKey",
  "fqn": "pulumi_gandi.domain",
  "classes": {
   "gandi:domain/dnssecKey:DnssecKey": "DnssecKey"
  }
 },
 {
  "pkg": "gandi",
  "mod": "domain/domain",
  "fqn": "pulumi_gandi.domain",
  "classes": {
   "gandi:domain/domain:Domain": "Domain"
  }
 },
 {
  "pkg": "gandi",
  "mod": "domain/glueRecord",
  "fqn": "pulumi_gandi.domain",
  "classes": {
   "gandi:domain/glueRecord:GlueRecord": "GlueRecord"
  }
 },
 {
  "pkg": "gandi",
  "mod": "domain/nameservers",
  "fqn": "pulumi_gandi.domain",
  "classes": {
   "gandi:domain/nameservers:Nameservers": "Nameservers"
  }
 },
 {
  "pkg": "gandi",
  "mod": "email/forwarding",
  "fqn": "pulumi_gandi.email",
  "classes": {
   "gandi:email/forwarding:Forwarding": "Forwarding"
  }
 },
 {
  "pkg": "gandi",
  "mod": "email/mailbox",
  "fqn": "pulumi_gandi.email",
  "classes": {
   "gandi:email/mailbox:Mailbox": "Mailbox"
  }
 },
 {
  "pkg": "gandi",
  "mod": "livedns/domain",
  "fqn": "pulumi_gandi.livedns",
  "classes": {
   "gandi:livedns/domain:Domain": "Domain"
  }
 },
 {
  "pkg": "gandi",
  "mod": "livedns/key",
  "fqn": "pulumi_gandi.livedns",
  "classes": {
   "gandi:livedns/key:Key": "Key"
  }
 },
 {
  "pkg": "gandi",
  "mod": "livedns/record",
  "fqn": "pulumi_gandi.livedns",
  "classes": {
   "gandi:livedns/record:Record": "Record"
  }
 },
 {
  "pkg": "gandi",
  "mod": "simplehosting/instance",
  "fqn": "pulumi_gandi.simplehosting",
  "classes": {
   "gandi:simplehosting/instance:Instance": "Instance"
  }
 },
 {
  "pkg": "gandi",
  "mod": "simplehosting/vhost",
  "fqn": "pulumi_gandi.simplehosting",
  "classes": {
   "gandi:simplehosting/vhost:Vhost": "Vhost"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "gandi",
  "token": "pulumi:providers:gandi",
  "fqn": "pulumi_gandi",
  "class": "Provider"
 }
]
"""
)
