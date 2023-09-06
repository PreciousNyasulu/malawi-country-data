-- public.constituencies definition

-- Drop table

-- DROP TABLE public.constituencies;

CREATE TABLE public.constituencies (
	id int4 NOT NULL,
	"name" varchar NULL,
	district_id int4 NOT NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT constituencies_pk PRIMARY KEY (id)
);


-- public.districts definition

-- Drop table

-- DROP TABLE public.districts;

CREATE TABLE public.districts (
	id int4 NOT NULL,
	"name" varchar NULL,
	code varchar NULL,
	region varchar NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT districts_pk PRIMARY KEY (id)
);


-- public.residential_areas definition

-- Drop table

-- DROP TABLE public.residential_areas;

CREATE TABLE public.residential_areas (
	id int4 NOT NULL,
	district_id int4 NULL,
	"name" varchar NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT residential_areas_pk PRIMARY KEY (id)
);


-- public.traditional_authorities definition

-- Drop table

-- DROP TABLE public.traditional_authorities;

CREATE TABLE public.traditional_authorities (
	id int4 NOT NULL,
	"name" varchar NULL,
	district_id int4 NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT traditional_authorities_pk PRIMARY KEY (id)
);


-- public.villages definition

-- Drop table

-- DROP TABLE public.villages;

CREATE TABLE public.villages (
	id int4 NOT NULL,
	village_name varchar NULL,
	district_id int4 NULL,
	CONSTRAINT villages_pk PRIMARY KEY (id)
);


-- public.wards definition

-- Drop table

-- DROP TABLE public.wards;

CREATE TABLE public.wards (
	id int4 NOT NULL,
	"name" varchar NULL,
	constituency_id int4 NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT wards_pk PRIMARY KEY (id)
);

INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (1,'Iponjola',1,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (2,'Kapoka',1,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (3,'Mahowe',2,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (4,'Nthalire',2,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (5,'Zambwe',3,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (6,'Yamba',3,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (7,'Nkhangwa',4,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (8,'Hanga',4,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (9,'Chisenga',5,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (10,'Wenya',5,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (11,'Songwe',6,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (12,'Kaporo',6,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (13,'Chilanga',7,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (14,'Rukuru',7,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (15,'Lupembe',8,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (16,'Mlale',8,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (17,'Nyungwe',9,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (18,'Zgeba',9,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (19,'Khwawa',10,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (20,'Uliwa',10,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (21,'Tcharo',11,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (22,'Mchenga',11,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (23,'Chinyolo/Mphompha',12,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (24,'Mayembe',12,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (25,'Hewe - Mwazisi',13,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (26,'Nkhamanga',13,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (27,'Henga-Phoka',14,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (28,'Phoka',14,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (29,'Usisya',15,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (30,'Chikwina',15,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (31,'Boma',16,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (32,'Kandoli',16,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (33,'Chitheka',17,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (34,'Tchesamu',17,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (35,'Kavuzi',18,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (36,'Mpamba',18,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (37,'Sanga',19,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (38,'Chintheche',19,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (39,'Tukombo',20,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (40,'Mbamba',20,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (41,'Likoma',21,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (42,'Chizumulu',21,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (43,'Lupaso Nkholongo',22,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (44,'Chibanja',22,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (45,'Katawa',22,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (46,'Mzilawaingwe',22,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (47,'Mchengautuwa West',22,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (48,'Chibavi East',22,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (49,'Jombo/Kaning''ina',22,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (50,'Zolozolo East',22,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (51,'Masasa',22,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (52,'Msongwe',22,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (53,'Luwinga',22,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (54,'Chiputula',22,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (55,'Mchengautuwa West',22,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (56,'Chibavi West',22,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (57,'Zolozolo West',22,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (58,'Kasito East',23,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (59,'Kasito West',23,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (60,'Njuyu',24,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (61,'Ekwendeni',24,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (62,'Kafukule',25,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (63,'Emcisweni',25,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (64,'Luviri',26,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (65,'Mabiri',26,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (66,'Euthini',27,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (67,'Mbalachanda',27,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (68,'Bulala',28,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (69,'Mzalangwe',28,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (70,'Mabilabo North',29,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (71,'Mabilabo South',29,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (72,'Manyamula',30,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (73,'Perekezi',30,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (74,'Kampingo Central',31,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (75,'Walula',31,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (76,'Engalaweni',32,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (77,'Kapopo',32,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (78,'Khonsolo North',33,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (79,'Khonsolo South',33,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (80,'Nthembwe',34,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (81,'Lodjwa',34,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (82,'Mpeni',35,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (83,'Mthabua',35,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (84,'Lisasadzi',36,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (85,'Lifupa',36,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (86,'Mpasadzi',37,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (87,'Matenje',37,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (88,'Chigodi',38,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (89,'Rusa',38,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (90,'Mziza',39,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (91,'Chibophi',39,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (92,'Kachokolo',40,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (93,'Kambira',40,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (94,'Chipala',41,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (95,'Mponda',41,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (96,'Belere',41,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (97,'Chankhanga',41,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (98,'Chimbuna',41,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (99,'Chitete',41,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (100,'Chithiba',41,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (101,'Chithiba',41,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (102,'Bunda',41,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (103,'Kapalankhwazi',41,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (104,'Kasalika',41,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (105,'Kabvunguti',41,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (106,'Ndonda',42,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (107,'Mbongozi',42,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (108,'Kasitu',43,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (109,'Nkhunga',43,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (110,'Liwaladzi',45,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (111,'Bua',44,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (112,'Mpondagaga',45,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (113,'Boma',45,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (114,'Linga-Sani',46,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (115,'Mpama',46,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (116,'Mtosa',47,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (117,'Kasangadzi',47,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (118,'Katete-Thumba',48,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (119,'Kalira-Masamba',48,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (120,'Bawala',49,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (121,'Masangano',49,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (122,'Kamlenje',50,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (123,'Chimbwadzi',50,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (124,'Langa',51,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (125,'Mtsiro',51,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (126,'Chiwere North',52,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (127,'Chiwere East',52,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (128,'Mukula East',53,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (129,'Mukula West',53,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (130,'Msakambewa East',54,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (131,'Msakambewa West',54,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (132,'Chakhaza South-West',55,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (133,'Chakhaza North-West',55,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (134,'Mponela',56,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (135,'Dzoole North',56,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (136,'Dzoole South',57,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (137,'Kayembe',57,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (138,'Chakhaza North',58,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (139,'Chakhaza South',58,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (140,'Liwadzi',59,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (141,'Lingadzi',59,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (142,'Maganga-Kuluunda',60,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (143,'Kalonga',60,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (144,'Ndindi',61,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (145,'Kambalame',61,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (146,'Kambwiri',62,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (147,'Pemba',62,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (148,'Chitala',63,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (149,'Lipimbi',63,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (150,'Mchemani',64,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (151,'Luweledzi',64,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (152,'Mponda',65,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (153,'Kapiri',65,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (154,'Mikundi',66,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (155,'Msachembe',66,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (156,'Mkoma',67,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (157,'Mtope',67,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (158,'Msitu',68,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (159,'Namnjiwa',68,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (160,'Kalumbe',69,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (161,'Chimimbe',69,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (162,'Mapuyu',70,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (163,'Nyanja',70,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (164,'Kachawa',71,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (165,'Kamanzi',71,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (166,'Demera',72,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (167,'Kalambe',72,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (168,'Mlodzenzi',73,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (169,'Dzanzi',73,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (170,'Nsanjiko',74,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (171,'Ngala',74,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (172,'Njewa',75,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (173,'Nalikule',76,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (174,'Chiwamba',76,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (175,'Nkhoma',77,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (176,'Chilenje',77,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (177,'Mkuza',78,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (178,'Mazengera',78,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (179,'Mtenthera',79,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (180,'Nyanja',79,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (181,'Chowo',80,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (182,'Chitsime',81,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (183,'Mlodza',81,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (184,'Chilobwe',82,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (185,'Nyang''amire',82,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (186,'Mteza',83,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (187,'Milindi',83,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (188,'Chiwenga',84,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (189,'Mthunthumala',84,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (190,'Bunda',85,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (191,'Katope',85,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (192,'Msinja',86,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (193,'Nadzumi',86,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (194,'Malingunde',87,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (195,'Chiputu',87,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (196,'Chisapo I',88,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (197,'Chisapo II',88,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (198,'Phwetekere',88,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (199,'Mbidzi',88,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (200,'Chigoneka',88,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (201,'Mtsiliza',88,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (202,'Mtandire',88,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (203,'Mvunguti',89,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (204,'M''gona - Chitata',89,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (205,'Kabwabwa',89,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (206,'Chimutu',89,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (207,'Nyama',89,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (208,'Maria',89,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (209,'Nankhaka',90,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (210,'Lumbadzi',90,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (211,'State House',91,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (212,'Chilinde I',91,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (213,'Chilinde II',91,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (214,'Kaliyeka',91,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (215,'Tsabango I',91,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (216,'Tsabango II',91,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (217,'Kawale I',92,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (218,'Kawale II',92,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (219,'Ngwenya',92,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (220,'Biwi',92,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (221,'Mchesi',92,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (222,'Sese',92,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (223,'Mayani North',93,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (224,'Mayani South',93,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (225,'Thiwi',94,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (226,'Diamphwe',94,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (227,'Malembo',95,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (228,'Thete',95,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (229,'Katewe',96,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (230,'Magomero',96,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (231,'Golomoti',97,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (232,'Mtakataka',97,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (233,'Bembeke',98,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (234,'Mlunduni',98,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (235,'Msunduzi',99,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (236,'Kambirimbiri',99,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (237,'Mkundi',100,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (238,'Chongoni',100,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (239,'Mphepozinai',101,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (240,'Kandeu',101,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (241,'Kasinje',102,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (242,'Sharpe Valley',102,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (243,'Mwalawoyera',103,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (244,'Bilira',103,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (245,'Dzunje',104,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (246,'Livilidzi',104,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (247,'Likudzi',105,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (248,'Ntonda',105,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (249,'Njolomole',106,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (250,'Lizulu',106,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (251,'Gomani Chikuse',107,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (252,'Tsangano',107,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (253,'Makanjira North',108,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (254,'Makanjira South',108,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (255,'Malindi',109,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (256,'Mikongo',109,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (257,'Maiwa',110,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (258,'Masanje',110,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (259,'Katuli North',111,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (260,'Katuli South',111,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (261,'Nlimba',112,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (262,'Chipunga',112,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (263,'Chilipa',113,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (264,'Katemba-Mtimabi',113,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (265,'Koche',114,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (266,'Thundu',114,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (267,'Mtumbwasi',115,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (268,'Mwasa',115,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (269,'Chigawe',115,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (270,'Chikole',115,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (271,'Kalungu',115,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (272,'Mikomwa',115,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (273,'Nkanamwano',115,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (274,'Msikisi',115,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (275,'Msukamwere',115,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (276,'Ndege',115,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (277,'Nkungulu',116,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (278,'Mpale',116,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (279,'Mvumba',117,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (280,'Malembo',117,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (281,'Monkey Bay',118,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (282,'Nkope',118,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (283,'Namavi',119,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (284,'Mbwazi',119,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (285,'Majuni',120,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (286,'Mandimba',120,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (287,'Shire',121,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (288,'Nkhonde',121,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (289,'Bwaila',122,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (290,'Liwawadzi',122,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (291,'Liviridzi',123,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (292,'Mulunguzi',123,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (293,'Utale',124,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (294,'Chimwalire',124,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (295,'Nyambi',125,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (296,'Mpiri',125,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (297,'Nsanama',126,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (298,'Mbonechera',126,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (299,'Mlomba',127,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (300,'Kanjuli',127,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (301,'Kawinga',128,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (302,'Nkoola',128,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (303,'Mposa',129,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (304,'Chamba',129,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (305,'Chinguni',130,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (306,'Lisanjala',130,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (307,'Chikweo',131,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (308,'Ngokwe',131,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (309,'Naisi',132,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (310,'Songani',132,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (311,'Chikomwe',133,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (312,'Chimwalira',133,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (313,'Linthipe',134,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (314,'Chingale',134,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (315,'Nswaswa',135,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (316,'Namilongo',135,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (317,'Mtungulutsi',136,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (318,'Chisenjere',136,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (319,'Lifani',137,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (320,'Naming''azi',137,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (321,'Chanda',138,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (322,'Ulumba',138,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (323,'Mbindi',139,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (324,'Chiphoola',139,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (325,'Chilwa',140,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (326,'Matiya',140,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (327,'Masongola',141,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (328,'Chinamwali',141,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (329,'Chirunga',141,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (330,'Central',141,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (331,'Likangala',141,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (332,'Chambo',141,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (333,'Sadzi',141,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (334,'Mbedza',141,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (335,'Mtiya',141,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (336,'Mpira',141,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (337,'Mitumbira',142,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (338,'Chikowa',142,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (339,'Mombezi',143,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (340,'Nguludi',143,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (341,'Nangalukutiche',144,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (342,'Ntayamwana',144,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (343,'Mwanje',145,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (344,'Thumbwe',145,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (345,'Mulombozi',146,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (346,'Lirangwe',146,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (347,'Chikwembere',147,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (348,'Linjidzi',147,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (349,'Matindi',148,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (350,'Lunzu',148,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (351,'Chilaweni',149,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (352,'Mudi',149,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (353,'Mpemba',150,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (354,'Soche',150,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (355,'Chilangoma',151,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (356,'Chikowa',151,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (357,'Namalimwe',152,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (358,'Ndirande Matope',152,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (359,'Ndirande Gamulani',153,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (360,'Ndirande Makata',153,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (361,'Nyambadwe',153,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (362,'Soche East',154,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (363,'Soche West',154,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (364,'Green Corner',154,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (365,'Blantyre South',154,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (366,'Naotcha',154,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (367,'Mapanga',155,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (368,'Nkolokoti',155,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (369,'Limbe Central',155,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (370,'Bangwe Nthandizi',156,'2019-02-12 07:06:16','2019-02-12 07:06:16');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (371,'Mzedi',156,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (372,'Bangwe',156,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (373,'Namiyango',157,'2019-02-12 07:06:16','2019-02-12 07:06:16'),
	 (374,'Chigumula',157,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (375,'Misesa',157,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (376,'Mzamba Nantipwiri',157,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (377,'Makungwa',157,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (378,'Blantyre City Central',158,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (379,'Chilomoni',158,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (380,'Chigwaja',158,'2019-02-12 07:06:17','2019-02-12 07:06:17');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (381,'Mbayani',159,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (382,'Michiru',159,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (383,'South Lunzu',159,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (384,'Khudze',160,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (385,'Mitseche',160,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (386,'Thambani',161,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (387,'Mpandadzi',161,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (388,'Lisungwi',162,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (389,'Ligowe',162,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (390,'Chilimbondo',163,'2019-02-12 07:06:17','2019-02-12 07:06:17');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (391,'Chikonde',163,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (392,'Makungwa',164,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (393,'Mikolongwe',164,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (394,'Masenjere',165,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (395,'Didi',165,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (396,'Khonjeni',166,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (397,'Nchima',166,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (398,'Thekerani',167,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (399,'Mapanga',167,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (400,'Magunda',168,'2019-02-12 07:06:17','2019-02-12 07:06:17');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (401,'Muonekera',168,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (402,'Namadzi',169,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (403,'Namisonga',169,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (404,'Sambagaru',169,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (405,'Lolo',169,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (406,'Mthundu',169,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (407,'Luchenza',169,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (408,'Mapanga',169,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (409,'Kapiri',169,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (410,'Masambanjati',170,'2019-02-12 07:06:17','2019-02-12 07:06:17');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (411,'Namagazi',170,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (412,'Thava',171,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (413,'Dzimbiri',171,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (414,'Mpasa',172,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (415,'Likulezi',172,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (416,'Machemba',173,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (417,'Migowi',173,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (418,'Khongoloni',174,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (419,'Thundu',174,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (420,'Chiringa',175,'2019-02-12 07:06:17','2019-02-12 07:06:17');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (421,'Sukasanje',175,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (422,'Mauzi',176,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (423,'Swang''oma',176,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (424,'Mimosa',177,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (425,'Milonde',177,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (426,'Chitakale',178,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (427,'Likhubula',178,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (428,'Chisitu',179,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (429,'Ntenjera',179,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (430,'Muloza',180,'2019-02-12 07:06:17','2019-02-12 07:06:17');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (431,'Limbuli',180,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (432,'Msikawanjala',181,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (433,'Mkumbiza',181,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (434,'Mulemba',182,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (435,'Chikuli',182,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (436,'Chole',183,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (437,'Chambe',183,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (438,'Nambilanje',184,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (439,'Namboya',184,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (440,'Mulomba',185,'2019-02-12 07:06:17','2019-02-12 07:06:17');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (441,'Mombezi',185,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (442,'Kawanda',186,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (443,'Nyakamba',186,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (444,'Mikalango',187,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (445,'Alumenda',187,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (446,'Lengwe',188,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (447,'Bwabwali',188,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (448,'Mwamphanzi',189,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (449,'Ndalanda',189,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (450,'Makhwira North',190,'2019-02-12 07:06:17','2019-02-12 07:06:17');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (451,'Makhwira South',190,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (452,'Chibisa',191,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (453,'Chimwanjale',191,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (454,'Nyachilenda',192,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (455,'Matundu',192,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (456,'Dinde',193,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (457,'Chekelere',193,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (458,'Chigumukire',194,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (459,'Misamvu',194,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (460,'Lalanje',195,'2019-02-12 07:06:17','2019-02-12 07:06:17');
INSERT INTO public.wards (id,name,constituency_id,created_at,updated_at) VALUES
	 (461,'Mlonda',195,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (462,'Kalulu',196,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (463,'Ruo',196,'2019-02-12 07:06:17','2019-02-12 07:06:17'),
	 (464,'Area 18',89,'2019-03-17 04:50:57','2019-03-17 04:50:57');