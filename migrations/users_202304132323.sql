INSERT INTO public.users (name,email,"password","role",provider,photo,verified,is_admin,created_at,updated_at,deleted_at) VALUES
	 ('Admin User','admin@example.com','$2a$10$5mtZ4/BqFlgGpFIdq7tc5eye0BtaUfud/GjTrxDjKkpU5ZHXUJK5O','user','local','imgAdmin.png',false,true,'2023-04-13 22:47:55.66424+02','2023-04-13 22:47:55.66424+02',NULL),
	 ('John Doe','john@example.com','$2a$10$vxvsI2y4tvLYocOgQP8p1ezdm2ky1WR5h.cNoHsgOfbBI...VykUS','user','local','imgJohn.png',false,false,'2023-04-13 22:49:02.795436+02','2023-04-13 22:49:02.795436+02',NULL),
	 ('Jane Doe','jane@example.com','$2a$10$PhdsHIIc.iGwPOhDwV4i2uBMd/Oz/ivbz0RbLZqlLr.e3txeotyAu','user','local','imgJohn.png',false,false,'2023-04-13 22:49:45.511377+02','2023-04-13 22:49:45.511377+02',NULL);
