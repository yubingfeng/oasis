

from openstack import connection


username = "yuyue_paas"
password = "seek123$Truth"
userDomainId = "yuyue_paas"
auth_url = "https://iam.cn-north-1.myhwcloud.com/v3"
project_id="fef67179ca37474da0c759a9f31d6c28"

conn = connection.Connection(auth_url=auth_url,
                             user_domain_id = userDomainId,
                             project_id=projectId,
                             username=username,
                             password=password)

print("the conn is ",conn)

