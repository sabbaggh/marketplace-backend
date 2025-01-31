CREATE TABLE usuarios (
	id_usuario INT auto_increment NOT NULL PRIMARY KEY,
    username                 VARCHAR(20) NOT NULL, 
    password                    VARCHAR(80) NOT NULL, 
    fecha_nacimiento DATE NOT NULL,
    tecnologia BOOL NOT NULL,
    ropa BOOL NOT NULL,
    perfumes BOOL NOT NULL,
    videojuegos BOOL NOT NULL,
    coleccionables BOOL NOT NULL,
    calificacion_comprador FLOAT DEFAULT 0,
	calificacion_vendedor FLOAT DEFAULT 0
);

CREATE TABLE publicacion (
	id_publicacion INT auto_increment NOT NULL PRIMARY KEY,
    id_usuario INT,
    titulo VARCHAR(40) NOT NULL,
    descripcion VARCHAR(200) NOT NULL,
    tipo_entrega VARCHAR(30) NOT NULL,
    categoria VARCHAR(50) NOT NULL,
    precio FLOAT NOT NULL,
    ubicacion VARCHAR(30),
    estatus ENUM('ACTIVA', 'ELIMINADA', 'ARCHIVADA') DEFAULT 'ACTIVA',
    FOREIGN KEY (id_usuario) REFERENCES usuarios(id_usuario)
);

CREATE TABLE ofertas (
	id_oferta INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    id_publicacion INT,
    id_comprador INT,
    precio FLOAT NOT NULL,
    estatus ENUM('PENDIENTE', 'ACEPTADA', 'RECHAZADA', 'CANCELADA') DEFAULT 'PENDIENTE',
    fecha_oferta TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_publicacion) REFERENCES publicacion(id_publicacion),
    FOREIGN KEY (id_comprador) REFERENCES usuarios(id_usuario)
);

CREATE TABLE transacciones (
    id_transaccion INT AUTO_INCREMENT PRIMARY KEY,
    id_oferta INT,
    id_vendedor INT,
    id_comprador INT,
    monto FLOAT NOT NULL,
    estatus ENUM('EN_HOLD', 'LIBERADO', 'DISPUTA', 'REEMBOLSADO') DEFAULT 'EN_HOLD',
    FOREIGN KEY (id_oferta) REFERENCES ofertas(id_oferta),
    FOREIGN KEY (id_vendedor) REFERENCES usuarios(id_usuario),
    FOREIGN KEY (id_comprador) REFERENCES usuarios(id_usuario)
);

CREATE TABLE imagenes_publicacion (
    id_imagen INT AUTO_INCREMENT PRIMARY KEY,
    id_publicacion INT,
    url VARCHAR(255) NOT NULL,
    FOREIGN KEY (id_publicacion) REFERENCES publicacion(id_publicacion)
);



