-- +goose Up
-- +goose StatementBegin
INSERT INTO
    tours (
        title,
        location,
        category,
        tags,
        description,
        nights_count,
        program,
        included,
        not_included,
        difficulty_level,
        comfort_level,
        tour_dates,
        important_info,
        head_media,
        program_media,
        acc_media,
        faq,
        rating,
        map
    )
VALUES
    (
        'Исторические сокровища Казани',
        'Казань',
        'Тур',
        ARRAY ['Пешеходный'],
        'Этот увлекательный 4-дневный пешеходный тур призван погрузить вас в атмосферу исторических сокровищ Казани, столицы Республики Татарстан. Вас ждет увлекательное путешествие по многонациональному городу, где сочетаются восточные и западные культуры, удивительные архитектурные памятники и традиции.
Этот тур включает в себя проживание в уютном отеле, завтраки каждый день, экскурсионное обслуживание, транспортное обслуживание включая поездку в Болгар, обед в Болгарии, а также трансферы от/до аэропорта. Личные расходы, дополнительные услуги и процесс оплаты за дополнительное место не включены в стоимость тура.',
        4,
        ARRAY [
    'День 1:
08:00 - Прибытие в аэропорт Казани.
09:00 - Трансфер в отель, заселение.
10:00 - 13:00 - Первое знакомство с городом: прогулка по центральным улицам, посещение площади Тысячелетия.
13:00 - 14:00 - Обед.
14:30 - 17:00 - Экскурсия по Казанскому Кремлю.
18:00 - Свободное время, ужин.',
    'День 2:
09:00 - Завтрак в отеле.
10:00 - 12:30 - Экскурсия по Старому городу: Кул-Шариф мечеть, музей "Казанский Кремль".
13:00 - 14:00 - Обед.
14:30 - 17:30 - Прогулка по набережной Казанки, посещение Ленинградской площади, торговых рядов, Татарского музея.
18:00 - Свободное время на ужин.',
    'День 3:
- 08:30 - Завтрак в отеле.
- 09:30 - 12:30 - Экскурсия в музей-заповедник "Казанский Исторический".
- 13:00 - 14:00 - Обед национальной кухни.
- 14:30 - 17:30 - Прогулка по Бауманской улице, посещение памятника Мусе Джаляля, торговых рядов.
- 18:00 - Свободное время.',
    'День 4:
07:30 - Завтрак в отеле, выселение.
08:30 - 10:00 - Трансфер в Болгар.
10:00 - 13:00 - Экскурсия по Болгарскому историко-архитектурному музею.
13:30 - 14:30 - Обед в национальном стиле.
15:00 - 17:00 - Прогулка по окрестностям Болгара.
17:30 - 19:00 - Возвращение в Казань, свободное время для последних покупок и ужина.
21:00 - Трансфер в аэропорт.'],
        ARRAY ['Проживание в отеле на 3 ночи.','Завтраки в отеле.','Экскурсионное обслуживание.','Транспортное обслуживание (в день экскурсии в Болгар).','Обед в Болгарии.','Трансферы от/до аэропорта.'],
        ARRAY ['Авиабилеты до и из Казани.','Услуги фотографа','Дополнительные экскурсии и развлечения, не указанные в программе.','Ужины и обеды, кроме указанных.','Дополнительные услуги в отеле.'],
        'Базовый',
        'Высокий',
        ARRAY [to_date('2024-04-10', 'YYYY-MM-DD'),to_date('2024-04-14', 'YYYY-MM-DD'),
to_date('2024-04-20', 'YYYY-MM-DD'),to_date('2024-04-24', 'YYYY-MM-DD')],
        'Подготовьте удобную обувь и одежду для прогулок и экскурсий на природе
Возьмите с собой солнцезащитные средства, головной убор, средства от комаров и других насекомых
Убедитесь, что у вас есть документы: паспорт, страховка, билеты
Заранее ознакомьтесь с программой тура и списком вещей, которые необходимо взять с собой

Как оформить дополнительное место:
Для оформления дополнительного места в туре необходимо связаться с менеджером туристической компании, указав количество дополнительных участников и предпочтения по размещению и питанию
Дополнительные места могут быть оформлены при наличии свободных мест в гостиницах и на транспорте, либо в случае возможности приобретения дополнительных билетов на экскурсии
Оплата за дополнительное место производится в соответствии с тарифами туроператора, дополнительные услуги могут быть оплачены как вместе с основным туром, так и отдельно.

Условия отмены:
В случае отмены тура до 7 дней до начала, возврат средств осуществляется без комиссии
В случае отмены менее чем за 7 дней до начала тура, возврат средств может быть произведен с учетом штрафных условий

ВАЖНО: при отсутствии медицинских документов или противопоказаний от врача, возврат средств не производится',
        ARRAY ['https://i.imgur.com/cagCisi.jpeg','https://i.imgur.com/AxfjRgf.jpeg','https://i.imgur.com/9aGpWGt.jpeg','https://i.imgur.com/DWv12cG.jpeg'],
        ARRAY ['https://i.imgur.com/fL8hXTw.jpeg','https://i.imgur.com/Lv4TdIg.jpeg','https://i.imgur.com/8GKLJIs.jpeg','https://i.imgur.com/xsurq0N.jpeg'],
        ARRAY ['https://i.imgur.com/7hxzv6L.jpeg','https://i.imgur.com/QnTxuqF.jpeg','https://i.imgur.com/4UtqX2s.jpeg','https://i.imgur.com/qEtJ4w1.jpeg'],
        'FAQ',
        9,
        'https://yandex.ru/map-widget/v1/?ll=86.218161%2C51.377574&mode=routes&routes%5BactiveComparisonMode%5D=auto&routes%5BignoreTravelModes%5D=bicycle%2Cscooter&rtext=53.346785%2C83.776856~49.376355%2C87.482662~53.346785%2C83.776856&rtt=comparison&ruri=ymapsbm1%3A%2F%2Fgeo%3Fdata%3DCgg1MzExNzk1MBI50KDQvtGB0YHQuNGPLCDQkNC70YLQsNC50YHQutC40Lkg0LrRgNCw0LksINCR0LDRgNC90LDRg9C7IgoNwY2nQhUcY1VC~ymapsbm1%3A%2F%2Forg%3Foid%3D116115703057~ymapsbm1%3A%2F%2Fgeo%3Fdata%3DCgg1MzExNzk1MBI50KDQvtGB0YHQuNGPLCDQkNC70YLQsNC50YHQutC40Lkg0LrRgNCw0LksINCR0LDRgNC90LDRg9C7IgoNwY2nQhUcY1VC&utm_medium=mapframe&utm_source=maps&z=6.99'
    );

INSERT INTO
    tours (
        title,
        location,
        category,
        tags,
        description,
        nights_count,
        program,
        included,
        not_included,
        difficulty_level,
        comfort_level,
        tour_dates,
        important_info,
        head_media,
        program_media,
        acc_media,
        faq,
        rating,
        map
    )
VALUES
    (
        'Сказочный Казанский Квартал',
        'Казань',
        'Тур',
        ARRAY ['Пешеходный', 'Экскурсионный'],
        'Отправляйтесь в увлекательное путешествие по сказочным улочкам и кварталам Казани с нашим эксклюзивным туром "Сказочный Казанский Кварталы"! Погрузитесь в захватывающий мир истории и культуры этого уникального города, окутанного магией прошлого.
Пройдитесь по старинным улицам, где каждый камень несет в себе легенды и сказки, посетите музеи, галереи и цветочные рынки, окунитесь в атмосферу изысканного искусства и народного творчества. Участвуйте в мастер-классе по созданию украшений в стиле татарского фольклора и попробуйте национальные блюда, олицетворяющие весь вкус и колорит этого удивительного региона.
Наши экскурсии наполнены удивительными открытиями и впечатляющими моментами, где вы сможете ощутить дух истории, слившийся с современностью. Предлагаем вам окунуться в мир сказок и чудес с туром "Сказочный Казанский Кварталы" и прочувствовать всю прелесть и таинственность этого уникального города!',
        3,
        ARRAY [
    'День 1:
9:00 - Встреча группы у отеля.
09:30 - 12:30 - Прогулка по историческому центру города, посещение старинных кварталов, известных своей архитектурой и узкими улочками.
13:00 - 14:00 - Обед в уютном кафе с местными деликатесами.
14:30 - 17:30 - Посещение Музея изобразительных искусств, где вы сможете насладиться произведениями русских и мировых художников.
18:00 - Свободное время для отдыха или индивидуальных прогулок по окрестностям.
20:00 - Ужин в ресторане с татарской кухней.',  
    'День 2:
08:30 - Завтрак в отеле.
09:30 - 11:30 - Экскурсия по старым улочкам, где вы услышите легенды и сказки, связанные с каждым местом.
12:00 - 13:00 - Мастер-класс по созданию украшений в национальном стиле.
13:30 - 14:30 - Обед с блюдами татарской кухни.
15:00 - 17:00 - Прогулка по цветочным рынкам города, где вы сможете насладиться ароматами и яркими красками.
18:00 - Свободное время на отдых или продолжение самостоятельных прогулок.
20:00 - Ужин в ресторане с видом на набережную Волги.'],
        ARRAY ['Проживание в отеле на 1 ночь.', 'Завтраки в отеле.', 'Обеды и ужины, как указано в программе.','Экскурсионное сопровождение.', 'Мастер-класс по созданию украшений.', 'Трансферы от/до аэропорта.'],
        ARRAY ['Дополнительные экскурсии и развлечения.', 'Страховка.', 'Авиабилеты до и из Казани.'],
        'Базовый',
        'Высокий',
        ARRAY [to_date('2024-04-13', 'YYYY-MM-DD'), to_date('2024-04-14', 'YYYY-MM-DD'),
to_date('2024-04-19', 'YYYY-MM-DD'), to_date('2024-04-20', 'YYYY-MM-DD')],
        'Подготовьте удобную обувь и одежду для прогулок и экскурсий на природе
Возьмите с собой солнцезащитные средства, головной убор, средства от комаров и других насекомых
Убедитесь, что у вас есть документы: паспорт, страховка, билеты
Заранее ознакомьтесь с программой тура и списком вещей, которые необходимо взять с собой

Как оформить дополнительное место:
Для оформления дополнительного места в туре необходимо связаться с менеджером туристической компании, указав количество дополнительных участников и предпочтения по размещению и питанию
Дополнительные места могут быть оформлены при наличии свободных мест в гостиницах и на транспорте, либо в случае возможности приобретения дополнительных билетов на экскурсии
Оплата за дополнительное место производится в соответствии с тарифами туроператора, дополнительные услуги могут быть оплачены как вместе с основным туром, так и отдельно.

Условия отмены:
В случае отмены тура до 7 дней до начала, возврат средств осуществляется без комиссии
В случае отмены менее чем за 7 дней до начала тура, возврат средств может быть произведен с учетом штрафных условий
ВАЖНО: при отсутствии медицинских документов или противопоказаний от врача, возврат средств не производится',
        ARRAY ['https://i.imgur.com/4XWZ8nJ.jpeg','https://i.imgur.com/cagCisi.jpeg','https://i.imgur.com/AxfjRgf.jpeg','https://i.imgur.com/1nUbOdk.jpeg'],
        ARRAY ['https://i.imgur.com/Yteu1Yn.jpeg','https://i.imgur.com/fL8hXTw.jpeg','https://i.imgur.com/hjtaOTP.jpeg','https://i.imgur.com/GyLSCGx.jpeg'],
        ARRAY ['https://i.imgur.com/U68uxGI.jpeg','https://i.imgur.com/FvVXlvP.jpeg','https://i.imgur.com/4UtqX2s.jpeg','https://i.imgur.com/qEtJ4w1.jpeg'],
        'FAQ',
        9,
        'https://yandex.ru/map-widget/v1/?ll=86.571061%2C51.098385&mode=routes&routes%5BactiveComparisonMode%5D=auto&routes%5BignoreTravelModes%5D=bicycle%2Cscooter&rtext=51.957804%2C85.960634~50.208052%2C87.933155~51.588669%2C87.665591~50.910297%2C88.216293~51.957804%2C85.960634&rtt=comparison&ruri=ymapsbm1%3A%2F%2Fgeo%3Fdata%3DCgg1MzExNTcxNxJI0KDQvtGB0YHQuNGPLCDQoNC10YHQv9GD0LHQu9C40LrQsCDQkNC70YLQsNC5LCDQk9C-0YDQvdC-LdCQ0LvRgtCw0LnRgdC6IgoN2OurQhXL1E9C~~~ymapsbm1%3A%2F%2Forg%3Foid%3D243504873227~ymapsbm1%3A%2F%2Fgeo%3Fdata%3DCgg1MzExNTcxNxJI0KDQvtGB0YHQuNGPLCDQoNC10YHQv9GD0LHQu9C40LrQsCDQkNC70YLQsNC5LCDQk9C-0YDQvdC-LdCQ0LvRgtCw0LnRgdC6IgoN2OurQhXL1E9C&utm_medium=mapframe&utm_source=maps&z=7.61'
    );

INSERT INTO
    tours (
        title,
        location,
        category,
        tags,
        description,
        nights_count,
        program,
        included,
        not_included,
        difficulty_level,
        comfort_level,
        tour_dates,
        important_info,
        head_media,
        program_media,
        acc_media,
        faq,
        rating,
        map
    )
VALUES
    (
        'Трекинг по Уралу',
        'Уральские горы',
        'Тур',
        ARRAY ['Природа', 'Пеший поход', 'Приключения'],
        'Пройдите увлекательный трек по живописным местам Уральских гор, наслаждаясь красотой дикой природы.',
        3,
        ARRAY ['День 1: Начало путешествия и ознакомление с маршрутом',  
      'День 2: Восхождение на вершину горы', 
      'День 3: Спуск в долину и завершение путешествия'],
        ARRAY ['Проживание в палатках', 'Трехразовое питание', 'Транспорт включен в стоимость'],
        ARRAY ['Авиабилеты', 'Личные расходы', 'Страховка'],
        'Средний',
        'Средний',
        ARRAY [to_date('2024-11-11','YYYY-MM-DD'), to_date('2024-11-13','YYYY-MM-DD')],
        'Пожалуйста, обеспечьте себя надежной спальником и подходящей одеждой.',
        ARRAY ['https://i.imgur.com/X6Bo5Uw.png', 'https://i.imgur.com/Ka94fpN.png','https://i.imgur.com/N1SYYIa.jpeg','https://i.imgur.com/S7HiiAi.png'],
        ARRAY ['https://i.imgur.com/n9FnYUl.png','https://i.imgur.com/5WacxQs.jpeg','https://i.imgur.com/h6KZNcI.png','https://i.imgur.com/XcIBbnF.png'],
        ARRAY ['https://i.imgur.com/S2JNOmN.png','https://i.imgur.com/CSWJQIO.png','https://i.imgur.com/YfY6lg5.png','https://i.imgur.com/uUMQz9C.png'],
        'FAQ',
        9,
        'https://yandex.ru/map-widget/v1/?ll=49.212300%2C55.702312&mode=routes&routes%5BactiveComparisonMode%5D=auto&rtext=55.608092%2C49.298596~55.795530%2C49.124813~55.608092%2C49.298596&rtt=comparison&ruri=ymapsbm1%3A%2F%2Forg%3Foid%3D1040136978~ymapsbm1%3A%2F%2Forg%3Foid%3D1063471493~ymapsbm1%3A%2F%2Forg%3Foid%3D1040136978&utm_medium=mapframe&utm_source=maps&z=11.28'
    );

INSERT INTO
    tours (
        title,
        location,
        category,
        tags,
        description,
        nights_count,
        program,
        included,
        not_included,
        difficulty_level,
        comfort_level,
        tour_dates,
        important_info,
        head_media,
        program_media,
        acc_media,
        faq,
        rating,
        map
    )
VALUES
    (
        'Культурный тур по Казани',
        'г. Казань',
        'Тур',
        ARRAY ['История', 'Архитектура', 'Пешый'],
        'Погрузитесь в богатую историю и культуру Казани, посетив ее исторические места и насладившись вкусом национальной кухни.',
        5,
        ARRAY ['День 1: Прибытие и ознакомительная экскурсия по городу',  
      'День 2: Посещение культурных памятников', 
      'День 3: Гастрономический тур по ресторанам', 
      'День 4: Экскурсия на историческую территорию Казанского Кремля', 
      'День 5: Свободное время и отъезд'],
        ARRAY ['Проживание в отеле', 'Завтраки', 'Транспорт включен в стоимость'],
        ARRAY ['Обеды и ужины', 'Личные расходы', 'Сувениры'],
        'Легкий',
        'Комфорт',
        ARRAY [to_date('2024-10-20', 'YYYY-MM-DD'),to_date('2024-10-25', 'YYYY-MM-DD'), to_date('2024-11-20', 'YYYY-MM-DD'), to_date('2024-11-25', 'YYYY-MM-DD')],
        'Не забудьте камеру и деньги на покупки сувениров.',
        ARRAY ['https://i.imgur.com/4imcw3V.png', 'https://i.imgur.com/Yw0LNbU.png','https://i.imgur.com/MBpYxd9.png','https://i.imgur.com/DcbfgrC.png'],
        ARRAY ['https://i.imgur.com/ox0aKf7.png', 'https://i.imgur.com/oe1k1fJ.png','https://i.imgur.com/fSaspYp.png', 'https://i.imgur.com/mPnmd9B.png'],
        ARRAY ['https://i.imgur.com/ak81lrB.png', 'https://i.imgur.com/eyD5xbT.png','https://i.imgur.com/dKRf45c.png','https://i.imgur.com/EcjhHLw.png'],
        'FAQ',
        4,
        'https://yandex.ru/map-widget/v1/?ll=49.540339%2C55.364556&mode=routes&routes%5BactiveComparisonMode%5D=auto&routes%5BignoreTravelModes%5D=bicycle%2Cscooter&rtext=55.608092%2C49.298596~55.798379%2C49.105238~55.799940%2C49.105863~54.982239%2C49.022628~55.796127%2C49.106414&rtt=comparison&ruri=ymapsbm1%3A%2F%2Forg%3Foid%3D1040136978~ymapsbm1%3A%2F%2Forg%3Foid%3D1009467292~ymapsbm1%3A%2F%2Forg%3Foid%3D1047417940~~ymapsbm1%3A%2F%2Fgeo%3Fdata%3DCgg1MzE2NjUwMRJD0KDQvtGB0YHQuNGPLCDQoNC10YHQv9GD0LHQu9C40LrQsCDQotCw0YLQsNGA0YHRgtCw0L0sINCa0LDQt9Cw0L3RjCIKDfhsREIVPC9fQg%2C%2C&utm_medium=mapframe&utm_source=maps&z=9.08'
    );

INSERT INTO
    tours (
        title,
        location,
        category,
        tags,
        description,
        nights_count,
        program,
        included,
        not_included,
        difficulty_level,
        comfort_level,
        tour_dates,
        important_info,
        head_media,
        program_media,
        acc_media,
        faq,
        rating,
        map
    )
VALUES
    (
        'Историко-культурный маршрут по Алтаю',
        'Алтай',
        'Тур',
        ARRAY ['Экскурсионный', 'Выездной'],
        '"Историко-культурный маршрут по Алтаю" - это уникальная возможность окунуться в мир древних традиций и красот природы одного из самых загадочных регионов России - Алтая. 
За 5 дней вы побываете в самых живописных местах региона: от загадочной Курайской степи, где вас ждут национальные праздники и мастер-классы местных ремесленников, до живописного Телецкого озера, о котором ходят легенды. 
Вы познакомитесь с удивительным миром алтайской культуры, посетите древние памятники и музеи, узнаете об алтайских обрядах и традициях. А экскурсия в долину Чулышман позволит вам понаблюдать за уникальной флорой и фауной, а также восхититься древними петроглифами.
Все это будет организовано комфортно и беззаботно: проживание в уютных гостиницах, вкусное питание, профессиональные гиды и удобный транспорт. Отправляйтесь в это увлекательное путешествие и погрузитесь в неповторимую атмосферу Алтая!',
        5,
        ARRAY ['День 1: Прибытие в Горно-Алтайск
9:00-10:00 Встреча группы в аэропорту/на вокзале
10:00-12:00 Трансфер в отель, размещение
14:00-18:00 Экскурсия по Горно-Алтайску с посещением музеев и исторических достопримечательностей
19:00-21:30 Ужин в национальном стиле',
'День 2: Курайская степь
10:30 - Поездка в Курайскую степь
11:45-15:25 Посещение национальных поселений, знакомство с традициями и ремеслами местных жителей
16:30-19:00 Национальные праздники и танцы
20:30 - Возвращение в город, отдых','День 3: Телецкое озеро
9:45 - Поездка в Телецкое озеро
11:00-13:45 Экскурсия на катерах по озеру, наслаждение живописными пейзажами
14:30-17:30 Посещение музеев и памятников природы
18:30 -  Ужин на берегу озера','День 4: Долина Чулышман
8:30 - Поездка в долину Чулышман
10:00-13:30 Посещение древних петроглифов, изучение истории исследований
14:00-17:00 Купание в горных озерах, пикник на природе
19:00 - Возвращение в город, досуг на выбор','День 5: Возвращение в Горно-Алтайск
9:00-18:00 Свободное время для покупок сувениров, посещение рынков
19:00 - Отъезд в аэропорт/на вокзал
20:00 - Прощание с гидом и памятные подарки для участников тура'],
        ARRAY ['Проживание в отелях на базе завтраков','Транспортные услуги включая трансферы и экскурсии','Услуги профессионального гида','Входные билеты в музеи и на экскурсии','Питание согласно программе тура'],
        ARRAY ['Алкогольные напитки','Услуги фотографа','Дополнительные экскурсии и развлечения'],
        'Базовый',
        'Высокий',
        ARRAY [to_date('2024-04-15', 'YYYY-MM-DD'),to_date('2024-04-20', 'YYYY-MM-DD')],
        'Подготовка к туру:
Подготовьте удобную обувь и одежду для прогулок и экскурсий на природе
Возьмите с собой солнцезащитные средства, головной убор, средства от комаров и других насекомых
Убедитесь, что у вас есть документы: паспорт, страховка, билеты
Заранее ознакомьтесь с программой тура и списком вещей, которые необходимо взять с собой

Как оформить дополнительное место:
Для оформления дополнительного места в туре необходимо связаться с менеджером туристической компании, указав количество дополнительных участников и предпочтения по размещению и питанию
Дополнительные места могут быть оформлены при наличии свободных мест в гостиницах и на транспорте, либо в случае возможности приобретения дополнительных билетов на экскурсии
Оплата за дополнительное место производится в соответствии с тарифами туроператора, дополнительные услуги могут быть оплачены как вместе с основным туром, так и отдельно.

Условия отмены:
В случае отмены тура до 7 дней до начала, возврат средств осуществляется без комиссии
В случае отмены менее чем за 7 дней до начала тура, возврат средств может быть произведен с учетом штрафных условий
ВАЖНО: при отсутствии медицинских документов или противопоказаний от врача, возврат средств не производится',
        ARRAY ['https://i.imgur.com/4XWZ8nJ.jpeg','https://i.imgur.com/wY8pprC.jpeg','https://i.imgur.com/duTesHi.jpeg','https://i.imgur.com/0V510qX.jpeg'],
        ARRAY ['https://i.imgur.com/qwIg4zS.jpeg','https://i.imgur.com/xsurq0N.jpeg','https://i.imgur.com/fL8hXTw.jpeg','https://i.imgur.com/8GKLJIs.jpeg'],
        ARRAY ['https://i.imgur.com/ZASHrQV.jpeg','https://i.imgur.com/4FU4wa0.jpeg','https://i.imgur.com/Utm0Bjl.jpeg','https://i.imgur.com/ca2DGyc.jpeg'],
        'FAQ',
        5,
        'https://yandex.ru/map-widget/v1/?ll=86.571061%2C51.098385&mode=routes&routes%5BactiveComparisonMode%5D=auto&routes%5BignoreTravelModes%5D=bicycle%2Cscooter&rtext=51.957804%2C85.960634~50.208052%2C87.933155~51.588669%2C87.665591~50.910297%2C88.216293~51.957804%2C85.960634&rtt=comparison&ruri=ymapsbm1%3A%2F%2Fgeo%3Fdata%3DCgg1MzExNTcxNxJI0KDQvtGB0YHQuNGPLCDQoNC10YHQv9GD0LHQu9C40LrQsCDQkNC70YLQsNC5LCDQk9C-0YDQvdC-LdCQ0LvRgtCw0LnRgdC6IgoN2OurQhXL1E9C~~~ymapsbm1%3A%2F%2Forg%3Foid%3D243504873227~ymapsbm1%3A%2F%2Fgeo%3Fdata%3DCgg1MzExNTcxNxJI0KDQvtGB0YHQuNGPLCDQoNC10YHQv9GD0LHQu9C40LrQsCDQkNC70YLQsNC5LCDQk9C-0YDQvdC-LdCQ0LvRgtCw0LnRgdC6IgoN2OurQhXL1E9C&utm_medium=mapframe&utm_source=maps&z=7.61'
    );

INSERT INTO
    tours (
        title,
        location,
        category,
        tags,
        description,
        nights_count,
        program,
        included,
        not_included,
        difficulty_level,
        comfort_level,
        tour_dates,
        important_info,
        head_media,
        program_media,
        acc_media,
        faq,
        rating,
        map
    )
VALUES
    (
        'Алтай: дикая природа и древние культуры',
        'Алтай',
        'Тур',
        ARRAY ['Экскурсионный', 'Автобусный'],
        'Алтай - это уникальный регион России, который сочетает в себе дикую природу и богатую историю и культуру местных народов. Этот тур даст вам возможность увидеть самые красивые и знаменитые достопримечательности Алтая, побывать в древних памятниках истории и культуры, попробовать вкусную национальную кухню и насладиться гостеприимством местных жителей.
В течение четырех дней вы совершите незабываемое путешествие по Алтаю, посетив Барнаул, горы Алтая, музеи "Катунь", "Алтайская природа", "Алтайская культура", "Алтайские горы", "Алтайские древности", "Алтайские народы", "Алтайские традиции". Вы увидите красоту и величие Алтайских гор, узнаете о богатой истории и культуре местных народов, почувствуете дух Алтая.
Не откладывайте свое путешествие в Алтай, забронируйте тур сейчас и получите незабываемые впечатления от этого волшебного региона! Мы гарантируем вам высокое качество обслуживания и максимальный комфорт в путешествии. Закажите тур прямо сейчас и получите скидку на следующее путешествие!',
        4,
        ARRAY [
    'День 1:

09:00 - Встреча туристов в аэропорту или вокзале города Барнаула, трансфер в гостиницу
10:00 - Экскурсия по Барнаулу: посещение музея истории Алтая, музея краеведения, музея литературы и искусства
13:00 - Обед в ресторане
14:00 - Трансфер в горы Алтая, посещение музея "Катунь"
16:00 - Прибытие в гостиницу, свободное время',
    'День 2:

09:00 - Завтрак в гостинице
10:00 - Экскурсия в музей "Алтайская природа": посещение музея природы, музея экологии, музея животных
13:00 - Обед в ресторане
14:00 - Экскурсия в музей "Алтайская культура": посещение музея истории, музея этнографии, музея религии
16:00 - Свободное время
18:00 - Возвращение в гостиницу',
    'День 3:

09:00 - Завтрак в гостинице
10:00 - Экскурсия в музей "Алтайские горы": посещение музея гор, музея альпинизма, музея туризма
13:00 - Обед в ресторане
14:00 - Экскурсия в музей "Алтайские древности": посещение музея археологии, музея палеонтологии, музея истории древних цивилизаций
16:00 - Свободное время
18:00 - Возвращение в гостиницу',
    'День 4:

09:00 - Завтрак в гостинице
10:00 - Экскурсия в музей "Алтайские народы": посещение музея татарской культуры, музея алтайской культуры, музея русской культуры
13:00 - Обед в ресторане
14:00 - Экскурсия в музей "Алтайские традиции": посещение музея народных промыслов, музея народных ремесел, музея народной музыки
16:00 - Трансфер в аэропорт или вокзал города Барнаула, отъезд домой'],
        ARRAY ['Трансферы по маршруту','Экскурсии по программе','Проживание в гостиницах (3 ночи)','Питание (4 завтрака, 4 обеда)','Услуги гида'],
        ARRAY ['Авиабилеты или билеты на поезд до Барнаула и обратно','Личные расходы (сувениры, напитки, дополнительные экскурсии и услуги, не предусмотренные программой тура)','Медицинская страховка','Дополнительные трансферы (например, трансферы в аэропорт или вокзал вне указанного времени)','Дополнительное проживание в гостинице (например, при досрочном прибытии или позднем отъезде)'],
        'Базовый',
        'Средний',
        ARRAY [to_date('2024-04-17', 'YYYY-MM-DD'),to_date('2024-04-25', 'YYYY-MM-DD')],
        'Необходимые документы:
Заграничный паспорт и виза (если требуется)
Медицинская страховка
Копии билетов, бронирования гостиницы и других необходимых документов
Контактная информация туроператора, гида и гостиницы
Медицинские противопоказания:
Убедитесь, что вы не имеете медицинских противопоказаний для участия в туре. Если у вас есть хронические заболевания или другие медицинские проблемы, уточните у туроператора, можно ли участвовать в туре и какие ограничения могут быть наложены.
Взять с собой необходимые медикаменты и убедиться, что у вас есть достаточное количество медикаментов на весь срок пребывания в стране.
Перенос тура на другие даты:
Если вы не можете участвовать в туре в назначенные даты, уточните у туроператора, можно ли перенести тур на другие даты. Узнайте, какие условия переноса тура и какие штрафные санкции могут быть наложены.
Если тур переносится на другие даты, убедитесь, что вы имеете все необходимые документы и медицинскую страховку на новые даты пребывания в стране.
Уведомите туроператора, гида и гостиницу о переносе тура и уточните, какие изменения могут быть внесены в программу тура и стоимость тура.',
        ARRAY ['https://i.imgur.com/AxfjRgf.jpeg','https://i.imgur.com/wY8pprC.jpeg','https://i.imgur.com/duTesHi.jpeg','https://i.imgur.com/0V510qX.jpeg'],
        ARRAY ['https://i.imgur.com/8GKLJIs.jpeg','https://i.imgur.com/tHgSibl.jpeg','https://i.imgur.com/fL8hXTw.jpeg','https://i.imgur.com/8GKLJIs.jpeg'],
        ARRAY ['https://i.imgur.com/ak81lrB.png','https://i.imgur.com/qEtJ4w1.jpeg','https://i.imgur.com/9ydwtzn.jpeg','https://i.imgur.com/ca2DGyc.jpeg'],
        'FAQ',
        8,
        'https://yandex.ru/map-widget/v1/?ll=86.218161%2C51.377574&mode=routes&routes%5BactiveComparisonMode%5D=auto&routes%5BignoreTravelModes%5D=bicycle%2Cscooter&rtext=53.346785%2C83.776856~49.376355%2C87.482662~53.346785%2C83.776856&rtt=comparison&ruri=ymapsbm1%3A%2F%2Fgeo%3Fdata%3DCgg1MzExNzk1MBI50KDQvtGB0YHQuNGPLCDQkNC70YLQsNC50YHQutC40Lkg0LrRgNCw0LksINCR0LDRgNC90LDRg9C7IgoNwY2nQhUcY1VC~ymapsbm1%3A%2F%2Forg%3Foid%3D116115703057~ymapsbm1%3A%2F%2Fgeo%3Fdata%3DCgg1MzExNzk1MBI50KDQvtGB0YHQuNGPLCDQkNC70YLQsNC50YHQutC40Lkg0LrRgNCw0LksINCR0LDRgNC90LDRg9C7IgoNwY2nQhUcY1VC&utm_medium=mapframe&utm_source=maps&z=6.99'
    );

INSERT INTO
    tours (
        title,
        location,
        category,
        tags,
        description,
        nights_count,
        program,
        included,
        not_included,
        difficulty_level,
        comfort_level,
        tour_dates,
        important_info,
        head_media,
        program_media,
        acc_media,
        faq,
        rating,
        map
    )
VALUES
    (
        'Адыгея: Путешествие в мир природы, истории и культуры',
        'Адыгея',
        'Тур',
        ARRAY ['Выездной', 'Экскурсионный'],
        'Адыгея: Путешествие в мир природы, истории и культуры" - это уникальная возможность познакомиться с одним из самых красивых и неизведанных регионов России. Это путешествие приведет вас в сердце Кавказа, где вы сможете насладиться великолепными горными пейзажами, узнать о богатой истории и культуре Адыгеи, попробовать вкусную национальную кухню и почувствовать гостеприимство местных жителей. Адыгея: Путешествие в мир природы, истории и культуры" - это уникальная возможность познакомиться с одним из самых красивых и неизведанных регионов России. 
Это путешествие приведет вас в сердце Кавказа, где вы сможете насладиться великолепными горными пейзажами, узнать о богатой истории и культуре Адыгеи, попробовать вкусную национальную кухню и почувствовать гостеприимство местных жителей.',
        /* nights count */
        3,
        ARRAY [
    'День 1:
08:00 Прибытие в Майкоп
09:00 Обзорная экскурсия по Майкопу: посещение музея истории Адыгеи, мемориала павшим в Великой Отечественной войне, кафедральной мечети
12:00 Трансфер в горное село Гузерипль
13:00 Экскурсия по Гузериплю: посещение музея истории Гузерипля, обзорная экскурсия по селу, прогулка по горным тропам
18:00 Ночевка в гостинице в Гузерипле',
    'День 2:
08:00 Трансфер в горное село Хакуринохабль
09:00 Экскурсия по Хакуринохаблю: посещение музея истории Хакуринохабля, обзорная экскурсия по селу, прогулка по горным тропам
12:00 Трансфер в горное село Каменномостский
13:00 Обзорная экскурсия по Каменномостскому: посещение музея истории Каменномостского, мемориала павшим в войне, кафедрального собора
18:00 Ночевка в гостинице в Каменномостском',
    'День 3:
08:00 Трансфер в горное село Льготная Поляна
09:00 Экскурсия по Льготной Поляне: посещение музея истории Льготной Поляны, обзорная экскурсия по селу, прогулка по горным тропам
12:00 Трансфер в Майкоп
13:00 Отдых и свободное время в Майкопе
18:00 Отъезд домой'],
        ARRAY ['Трансферы', 'Проживание в отеле','Питание: завтраки, обеды, ужины', 'Экскурсии', 'Услуги гида'],
        ARRAY ['Авиабилеты', 'Страховка', 'Дополнительные экскурсии и услуги', 'Личные расходы'],
        'Средний',
        'Простой',
        ARRAY [to_date('2024-07-01', 'YYYY-MM-DD'),to_date('2024-07-4', 'YYYY-MM-DD'),to_date('2024-07-15', 'YYYY-MM-DD'),to_date('2024-07-22', 'YYYY-MM-DD'),to_date('2024-08-01', 'YYYY-MM-DD'),to_date('2024-08-08', 'YYYY-MM-DD'),to_date('2024-08-15', 'YYYY-MM-DD'),to_date('2024-08-22', 'YYYY-MM-DD')],
        'Нужна ли страховка для участия в туре? Да, рекомендуется оформить страховку на случай непредвиденных обстоятельств.
    Как оформить бронь на тур? Для бронирования тура необходимо заполнить форму на сайте или связаться с нашими менеджерами по телефону.
    Как оплатить тур? Оплата тура производится на сайте или банковским переводом.
    Какие документы необходимы для поездки? Для поездки необходим паспорт гражданина РФ.
    Как добраться до места начала тура? Тур начинается в городе Краснодаре, добраться до которого можно на самолете или на поезде. Мы предоставляем трансфер от аэропорта или вокзала до отеля.',
        ARRAY ['https://i.imgur.com/rjfgjLd.jpeg','https://i.imgur.com/0V510qX.jpeg','https://i.imgur.com/0V510qX.jpeg'],
        ARRAY ['https://i.imgur.com/Lv4TdIg.jpeg','https://i.imgur.com/tHgSibl.jpeg','https://i.imgur.com/Yteu1Yn.jpeg','https://i.imgur.com/8GKLJIs.jpeg'],
        ARRAY ['https://i.imgur.com/5J6CT2a.jpeg','https://i.imgur.com/G1kpesM.jpeg','https://i.imgur.com/9ydwtzn.jpeg','https://i.imgur.com/ca2DGyc.jpeg'],
        'FAQ',
        9,
        'https://yandex.ru/map-widget/v1/?ll=40.135557%2C44.507393&mode=routes&routes%5BactiveComparisonMode%5D=auto&routes%5BignoreTravelModes%5D=bicycle%2Cscooter&rtext=44.606683%2C40.105852~43.993633%2C40.132774~45.019271%2C40.236350~44.296028%2C40.183475~44.595985%2C40.043949~44.606683%2C40.105852&rtt=comparison&ruri=ymapsbm1%3A%2F%2Fgeo%3Fdata%3DCgg1MzE2NTk0ORI90KDQvtGB0YHQuNGPLCDQoNC10YHQv9GD0LHQu9C40LrQsCDQkNC00YvQs9C10Y8sINCc0LDQudC60L7QvyIKDWVsIEIVPm0yQg%2C%2C~~~~~ymapsbm1%3A%2F%2Fgeo%3Fdata%3DCgg1MzE2NTk0ORI90KDQvtGB0YHQuNGPLCDQoNC10YHQv9GD0LHQu9C40LrQsCDQkNC00YvQs9C10Y8sINCc0LDQudC60L7QvyIKDWVsIEIVPm0yQg%2C%2C&z=9.18'
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists tours;

-- +goose StatementEnd