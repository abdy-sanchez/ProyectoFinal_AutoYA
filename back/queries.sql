CREATE TABLE carros (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	carro_id int8 NOT NULL,
	estado_reserva bool DEFAULT false NULL,
	precio int8 NOT NULL,
	modelo_year int8 NOT NULL,
	modelo_nombre text NOT NULL,
	marca text NOT NULL,
	transmision text NOT NULL,
	combustible text NOT NULL,
	tipo text NOT NULL,
	puertas int8 NOT NULL,
	abs bool NOT NULL,
	img text NOT NULL,
	CONSTRAINT carros_pkey PRIMARY KEY (id)
);

CREATE INDEX idx_carros_deleted_at ON public.carros USING btree (deleted_at);

insert into carros (carro_id, estado_reserva, precio, modelo_year, modelo_nombre, marca, transmision, combustible, tipo, puertas, abs, img) 
values 
(1,false,3100,2013,'Zoe','Renault','Automatico','Electrico','Hatchback',5,true,'https://cdn.autobild.es/sites/navi.axelspringer.es/public/media/image/2013/03/307153-renault-zoe-electrico-delantera.jpg'),
(2,false,2900,2013,'Spark','Chevrolet','Manual','Gasolina','Hatchback',5,true,'https://static.motor.es/fotos-ficha/chevrolet/chevrolet-spark-7504_1.jpg'),
(3,false,5200,2010,'Pathfinder','Nissan','Manual','Gasolina','Todoterreno',5,true,'https://www.nissan.com.co/content/dam/Nissan/co/vehicles/pathfinder/my23/vlp/overview/nissan-2023-pathfinder_exclusive_4wd-rojo_bi_tono-negro-exterior-overview-arbolexterior-d.jpg'),
(4,false,15000,2024,'Fortuner','TOYOTA','Manual','Gasolina','Todoterreno',5,true,'https://www.toyota.com.co/wp-content/uploads/2020/08/super-blanco.png'),
(5,false,1500,2013,'Twingo','Renault','Manual','Gasolina','Coupe',3,true,'https://revistacarro.com.br/wp-content/uploads/2018/08/RENAULTTWINGO1-1024x614.jpg');

