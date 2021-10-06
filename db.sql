insert  into casbin_rule values
                                 (1,'p','anonymous','/v1/login','create','*'),
                                 (2,'p','SUPER-ADMIN|SYSTEM-CLERK|COMPANY-ADMIN|COMPANY-CLERK','/v1/users','read','*'),
                                 (3,'p','SUPER-ADMIN|SYSTEM-CLERK|COMPANY-ADMIN|COMPANY-CLERK','/v1/users/:user-id','read','*'),
                                 (4,'p','SUPER-ADMIN','/v1/users','create','*'),
                                 (5,'p','SUPER-ADMIN','/v1/users/:user-id','delete','*'),
                                 (6,'p','SUPER-ADMIN','/v1/users/:user-id','update','*'),

                                 (7,'p','SUPER-ADMIN|COMPANY-ADMIN','/v1/companies/:company-id/users','read','c11f8819-afe6-4367-bd8f-4d5adb553433'),
                                 (8,'p','SUPER-ADMIN|COMPANY-ADMIN','/v1/companies/:company-id/users/:user-id','read','c11f8819-afe6-4367-bd8f-4d5adb553433'),
                                 (9,'p','SUPER-ADMIN | COMPANY-ADMIN','/v1/companies/:company-id/users/:user-id','delete','c11f8819-afe6-4367-bd8f-4d5adb553433'),
                                 (10,'p','SUPER-ADMIN | COMPANY-ADMIN','/v1/companies/:company-id/add-users','create','c11f8819-afe6-4367-bd8f-4d5adb553433'),

                                 (11,'p','SUPER-ADMIN','/v1/companies','create','*'),
                                 (12,'p','SUPER-ADMIN|SYSTEM-CLERK|COMPANY-ADMIN|COMPANY-CLERK','/v1/companies','read','*'),
                                 (13,'p','SUPER-ADMIN|SYSTEM-CLERK|COMPANY-ADMIN|COMPANY-CLERK','/v1/companies/:company-id','read','*'),
                                 (14,'p','SUPER-ADMIN','/companies/:company-id','update','*'),
                                 (15,'p','SUPER-ADMIN','/companies/:company-id','delete','*'),

                                 ( 16,'p','SUPER-ADMIN|SYSTEM-CLERK|COMPANY-ADMIN|COMPANY-CLERK','/v1/notifications','create','*'),
                                 (17,'p','SUPER-ADMIN|SYSTEM-CLERK|COMPANY-ADMIN|COMPANY-CLERK','/v1/notifications','read','*'),
                                 (18,'p','SUPER-ADMIN|SYSTEM-CLERK|COMPANY-ADMIN|COMPANY-CLERK','/v1/notifications/:id','delete','*'),
                                 (19,'p','SUPER-ADMIN|SYSTEM-CLERK|COMPANY-ADMIN|COMPANY-CLERK','/v1/notifications/unread/publish','read','*'),
                                 (20,'p','SUPER-ADMIN|SYSTEM-CLERK|COMPANY-ADMIN|COMPANY-CLERK','/v1/notifications/unread/email','read','*'),
                                 (21,'p','SUPER-ADMIN|SYSTEM-CLERK|COMPANY-ADMIN|COMPANY-CLERK','/v1/notifications/sms','create','*'),
                                 (22,'p','SUPER-ADMIN|SYSTEM-CLERK|COMPANY-ADMIN|COMPANY-CLERK','/v1/notifications/unread/sms','read','*'),
                                 (23,'p','SUPER-ADMIN|SYSTEM-CLERK|COMPANY-ADMIN|COMPANY-CLERK','/v1/notifications/email','create','*'),
                                 (24,'p','SUPER-ADMIN|SYSTEM-CLERK|COMPANY-ADMIN|COMPANY-CLERK','/v1/notifications/unread/email','read','*'),
                                 (25,'p','SUPER-ADMIN','/v1/policies','create','*'),
                                 (26,'p','SUPER-ADMIN | COMPANY-ADMIN','/v1/policies','read','*'),
                                 (27,'p','SUPER-ADMIN | COMPANY-ADMIN','/v1/policies','update','*'),
                                 (28,'p','SUPER-ADMIN | COMPANY-ADMIN','/v1/policies','delete','*');