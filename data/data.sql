DELETE FROM accounts;
DELETE FROM accounts_tokens;
DELETE FROM replenishment;

ALTER SEQUENCE accounts_id_seq RESTART WITH 1;
ALTER SEQUENCE replenishment_id_seq RESTART WITH 1;

INSERT INTO accounts ( phone, balance, password ) VALUES ('+992000000001', 5.00, '$2a$10$OaUtjCNv2DT5x/dXcV.P3eYkIPIRtBr/v8Nluwifz6brSkfyXOh6m');
INSERT INTO accounts ( phone, balance, password,  identified) VALUES ('+992000000002', 50.00, '$2a$10$OaUtjCNv2DT5x/dXcV.P3eYkIPIRtBr/v8Nluwifz6brSkfyXOh6m', 'true');

INSERT INTO accounts_tokens (token, accounts_id, expire, created) VALUES ('466a932ba9637beb93999489a593961c725b113032be1a2d109e25111fb3d9f37d9e63ffbe592f3e2a9960a7abed7fdf7452058ae5db75fe79e635fc8161a37c7c7aad29d58b12e5318d5308c8c2b575ad6fe811d85bcb2ad06d4d3ce2bfd77f3a493c137c9929b3f8f9cd8348dce74bd0317ae2835efd916deaa075c21e6c7b6f286fb9381351b415222bf4e4f18e2566eca86c0e08f320dbd1bd5d0ca0181a835f3b96523f497bfa3d03353421a487ba932ae1cdb3368ac240b6f93119fedd48a08ff0ec09d5036de4c480da005b46b3b1d64c017651ba70edc8c87488fa7c32323737aa7882188b528db9a5a15f4dde18ebda58cc00abb27157dba4424a4f', 1, '2021-11-21 20:50:57.627873', '2021-11-21 19:50:57.627873');
INSERT INTO accounts_tokens (token, accounts_id, expire, created) VALUES ('473a5fe51cf0d98d28f2f2401211f05805a1fe0a5840b3292ed0e79773b3d65b26b1586d9bd82551b697ea4ba6da01e55da549948567632d01a96a3eeb43f5e83199d5112f0d659823d4a88a4b65abefbe04824a5155b0484fef293fb9c4cd92477e068e63fe403c9481f5c5ba3b6952884e681ac940cc766da033ff43c088232aa42efb8beab8c485ff91269465e3b0bc8277f51b6cd107738c08b4660fb3e96189b2f3be8a9aa39241cc344362beb977613d16d6b38df3caee4c2448c0d147b5f7ac4cf6b43a7b65e0f855e09408aec7b1d27cb28dd074004d899f0b7c38540e2708f497a1dc487aafc28caef76003b6366808fbb187024da96fdb2523354f', 2, '2021-11-21 20:50:57.627873', '2021-11-21 19:50:57.627873');

INSERT INTO replenishment ( accounts_id, ammount, created) VALUES ( 1, 1.00, '2021-11-22 00:11:10.320096');
INSERT INTO replenishment ( accounts_id, ammount, created) VALUES ( 1, 1.00, '2021-11-22 00:11:11.362156');
INSERT INTO replenishment ( accounts_id, ammount, created) VALUES ( 1, 1.00, '2021-11-22 00:11:12.415216');
INSERT INTO replenishment ( accounts_id, ammount, created) VALUES ( 1, 1.00, '2021-11-22 00:11:14.964362');
INSERT INTO replenishment ( accounts_id, ammount, created) VALUES ( 1, 1.00, '2021-11-22 00:11:14.964362');
INSERT INTO replenishment ( accounts_id, ammount, created) VALUES ( 2, 10.00, '2021-11-22 00:11:16.084426');
INSERT INTO replenishment ( accounts_id, ammount, created) VALUES ( 2, 10.00, '2021-11-22 00:11:17.290495');
INSERT INTO replenishment ( accounts_id, ammount, created) VALUES ( 2, 10.00, '2021-11-22 00:11:18.538566');
INSERT INTO replenishment ( accounts_id, ammount, created) VALUES ( 2, 10.00, '2021-11-22 00:11:19.710634');
INSERT INTO replenishment ( accounts_id, ammount, created) VALUES ( 2, 10.00, '2021-11-22 00:11:21.21872');
