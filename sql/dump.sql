DROP TABLE IF EXISTS `persons`;

CREATE TABLE `persons` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `age` int DEFAULT NULL,
  `contact` varchar(45) DEFAULT NULL,
  `gender` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 101 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

INSERT INTO
  `persons`
VALUES
  (1, 'Reba Paucek', 5, '430-231-0937 x1120', 'Female'),
(
    2,
    'Mr. Luis O\"Kon',
    55,
    '(230) 286-0594 x790',
    'Female'
  ),
(3, 'Timmy Corkery', 56, '+1 (930) 575-1533', 'Male'),
(4, 'Merlin Langosh V', 8, '(200) 435-5234', 'Female'),
(5, 'Richie Hand', 72, '200.938.4725 x0436', 'Female'),
(6, 'Fiona O\"Keefe', 76, '470-782-3658 x345', 'Male'),
(
    7,
    'Mr. Sigrid Schamberger',
    45,
    '+1 (660) 941-1754',
    'Male'
  ),
(
    8,
    'Ms. Savanna Jakubowski',
    29,
    '+16305512026',
    'Female'
  ),
(9, 'Abbie Hickle', 64, '980-979-0400', 'Female'),
(
    10,
    'Trace Goodwin',
    41,
    '620.695.4451 x598',
    'Female'
  ),
(
    11,
    'Ms. Claudine Kuhn IV',
    16,
    '320.788.2938 x6456',
    'Male'
  ),
(
    12,
    'Mr. Orlando Mertz III',
    36,
    '300.500.5767 x52521',
    'Male'
  ),
(
    13,
    'Ms. Phyllis Denesik III',
    10,
    '650-642-1711 x385',
    'Female'
  ),
(
    14,
    'Ms. Zoie Frami III',
    74,
    '+1-900-802-5640',
    'Male'
  ),
(
    15,
    'Luther King I',
    33,
    '980-910-9285 x632',
    'Female'
  ),
(
    16,
    'Adalberto Marquardt',
    50,
    '290-448-6947 x3740',
    'Female'
  ),
(
    17,
    'Alejandrin Rau',
    19,
    '(770) 579-9412 x9215',
    'Female'
  ),
(
    18,
    'Mr. Henderson Kiehn',
    67,
    '870-896-1196 x156',
    'Female'
  ),
(19, 'Reta Paucek', 47, '1-490-246-0479', 'Female'),
(
    20,
    'Emmett Cruickshank',
    84,
    '1-220-349-3812',
    'Male'
  ),
(
    21,
    'Mozell Dickinson',
    53,
    '+1-230-836-8839',
    'Male'
  ),
(
    22,
    'Ahmad Reichel',
    53,
    '570.645.9303 x1767',
    'Male'
  ),
(
    23,
    'Richie Gleichner',
    35,
    '+1-490-265-1811',
    'Female'
  ),
(24, 'Mr. Russ Weimann', 25, '930-549-6890', 'Female'),
(
    25,
    'Trevion Sanford',
    78,
    '1-290-534-6479',
    'Female'
  ),
(
    26,
    'Savannah Stamm',
    40,
    '1-820-995-1137 x396',
    'Male'
  ),
(27, 'Marlene Parker', 7, '590.883.1788', 'Female'),
(28, 'Arely Pagac DVM', 46, '+1.500.800.9714', 'Male'),
(29, 'Wilton Windler', 87, '(520) 965-8854', 'Female'),
(30, 'Jett Lockman V', 86, '+12408063647', 'Male'),
(
    31,
    'Orlo Konopelski',
    12,
    '+1 (330) 744-5498',
    'Male'
  ),
(
    32,
    'Ms. Ebony Bradtke IV',
    38,
    '1-650-550-1904 x3238',
    'Female'
  ),
(
    33,
    'Ms. Zora Hegmann',
    5,
    '1-970-809-6643 x98727',
    'Male'
  ),
(
    34,
    'Gideon Powlowski',
    90,
    '1-540-510-0194',
    'Female'
  ),
(
    35,
    'Ms. Heaven Runte IV',
    4,
    '750-353-5332',
    'Female'
  ),
(
    36,
    'Mr. Dusty Flatley MD',
    24,
    '590.970.4700 x373',
    'Female'
  ),
(37, 'Emmie Effertz', 16, '370.410.3103', 'Female'),
(
    38,
    'Meagan Lebsack',
    95,
    '+1-930-595-1413',
    'Female'
  ),
(
    39,
    'Bailee Ruecker',
    47,
    '490.587.5984 x5410',
    'Female'
  ),
(40, 'Carmine Walsh', 7, '+1.730.789.0747', 'Female'),
(
    41,
    'Ms. Roxane Wolff IV',
    81,
    '(860) 827-4756 x8354',
    'Male'
  ),
(42, 'Deron McGlynn', 44, '340.542.1901', 'Male'),
(43, 'Roderick Ledner', 28, '1-270-953-5420', 'Male'),
(44, 'Letha Pollich', 79, '1-810-310-7944', 'Male'),
(
    45,
    'Stanton Dibbert',
    59,
    '680-851-9735 x638',
    'Female'
  ),
(46, 'Royce West MD', 9, '580-634-8059', 'Female'),
(
    47,
    'Anika Bauch V',
    63,
    '1-770-837-0979 x2940',
    'Male'
  ),
(
    48,
    'Darrion Gulgowski',
    33,
    '960.696.9890 x0937',
    'Male'
  ),
(
    49,
    'Ms. Alyson Miller',
    26,
    '+1-550-685-4579',
    'Male'
  ),
(
    50,
    'Margaretta Barrows',
    78,
    '1-520-761-8260 x6719',
    'Male'
  ),
(51, 'Ernest Johnson', 94, '680-910-1553', 'Female'),
(52, 'Jany Kuvalis', 96, '830.534.3592', 'Male'),
(
    53,
    'Antwan Nitzsche',
    57,
    '250.870.5191 x813',
    'Female'
  ),
(
    54,
    'Justine Glover',
    49,
    '1-460-384-3160 x04352',
    'Female'
  ),
(55, 'Kiera Batz', 4, '920-234-1105', 'Male'),
(
    56,
    'Dayna Champlin',
    63,
    '+1-500-359-8053',
    'Female'
  ),
(57, 'Shane Howe', 89, '+1-340-553-2420', 'Male'),
(
    58,
    'Mr. Cedrick Ledner',
    70,
    '780-410-8353 x74569',
    'Male'
  ),
(
    59,
    'Ms. Sadye Brekke',
    95,
    '+1-890-643-2146',
    'Male'
  ),
(60, 'Price Wiza', 9, '+1-300-463-4803', 'Male'),
(
    61,
    'Mr. Harmon Windler',
    84,
    '(270) 210-6397',
    'Male'
  ),
(
    62,
    'Mr. Kayley Kautzer',
    33,
    '+1.850.676.9349',
    'Male'
  ),
(
    63,
    'Bernadette Schaefer',
    58,
    '1-620-942-9143 x56828',
    'Male'
  ),
(
    64,
    'Cordia Predovic',
    55,
    '840.329.6570 x0691',
    'Female'
  ),
(
    65,
    'Mafalda Stehr',
    30,
    '1-470-726-5635 x307',
    'Male'
  ),
(
    66,
    'Ryley Carroll',
    55,
    '1-900-333-4044 x5007',
    'Male'
  ),
(
    67,
    'Mr. Helmer Stroman',
    66,
    '210.930.6428',
    'Female'
  ),
(68, 'Price Graham', 70, '(610) 603-7915', 'Male'),
(69, 'Anita Hintz', 26, '(880) 910-2715 x610', 'Male'),
(70, 'Cristina DuBuque', 17, '330-710-1009', 'Male'),
(71, 'Gregg Dicki', 99, '(790) 336-0103 x128', 'Male'),
(
    72,
    'Ms. Emmy Willms',
    6,
    '760-502-5892 x5503',
    'Male'
  ),
(73, 'Zachery Kohler', 39, '+17708106353', 'Female'),
(
    74,
    'Sigmund Towne',
    24,
    '(600) 370-7239 x1222',
    'Male'
  ),
(75, 'Trenton Borer', 37, '230-569-1918', 'Male'),
(76, 'Whitney Marks', 56, '1-860-566-6295', 'Female'),
(77, 'Reed Reilly', 44, '580-886-2524', 'Male'),
(
    78,
    'Myrtle Koch MD',
    24,
    '1-510-534-9346 x1816',
    'Female'
  ),
(
    79,
    'Mr. Davin Considine',
    57,
    '1-550-249-8558',
    'Female'
  ),
(80, 'Phoebe Zieme', 79, '1-500-487-4003', 'Female'),
(
    81,
    'Javon Ullrich',
    50,
    '570.700.7328 x3766',
    'Male'
  ),
(82, 'Simeon Upton', 64, '370-833-6894 x2084', 'Male'),
(
    83,
    'Bret Stark',
    42,
    '410.958.1681 x58951',
    'Female'
  ),
(
    84,
    'Nadia Hamill',
    55,
    '990-343-9013 x91414',
    'Female'
  ),
(
    85,
    'Christine Romaguera III',
    33,
    '950-697-6168',
    'Female'
  ),
(86, 'Kyra Roberts', 84, '420.861.8342', 'Male'),
(
    87,
    'Ms. Marisa Davis',
    83,
    '+1 (930) 335-0211',
    'Male'
  ),
(
    88,
    'Ms. Marilie Ernser',
    51,
    '660.442.5244 x9420',
    'Male'
  ),
(
    89,
    'Ms. Lorna Reilly MD',
    81,
    '450-710-6080 x47846',
    'Female'
  ),
(
    90,
    'Quincy Rippin DVM',
    40,
    '(760) 370-8795 x41537',
    'Female'
  ),
(
    91,
    'Ms. Maryse Kohler MD',
    24,
    '580.459.0631 x076',
    'Female'
  ),
(
    92,
    'Orrin Osinski',
    51,
    '1-420-704-2785 x701',
    'Female'
  ),
(
    93,
    'Betty Schneider',
    38,
    '1-930-338-7356 x04433',
    'Male'
  ),
(
    94,
    'Elliot Morar',
    39,
    '1-300-910-7638 x5533',
    'Male'
  ),
(
    95,
    'Ms. Demetris Auer PhD',
    86,
    '380.835.1872 x1666',
    'Male'
  ),
(
    96,
    'Madaline Bechtelar DVM',
    46,
    '530-365-5674',
    'Male'
  ),
(
    97,
    'Genevieve Kuphal',
    88,
    '680-658-8438 x9367',
    'Female'
  ),
(
    98,
    'Daphnee Leannon',
    34,
    '+1 (460) 992-9647',
    'Female'
  ),
(
    99,
    'Edd Kirlin III',
    29,
    '+1-290-610-3332',
    'Female'
  ),
(
    100,
    'Ms. Marianna Conn DDS',
    64,
    '(290) 775-7289',
    'Female'
  );
