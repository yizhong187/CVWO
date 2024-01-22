--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1
-- Dumped by pg_dump version 16.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Data for Name: subforums; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.subforums (id, name, description, created_at, updated_at, photo_url) FROM stdin;
7	Sassy Pickle Serenade	A quirky country band from Nashville, their music focuses on the humorous side of farm life, with a special love for pickles.	2024-01-12 03:35:31.691624+08	2024-01-13 02:39:37.598409+08	https://m.media-amazon.com/images/I/51r4m3CNzlL._AC_UF1000,1000_QL80_.jpg
8	Galactic Kazoo Orchestra	A sci-fi themed experimental band from Tokyo, blending traditional instruments with kazoos to create a unique interstellar sound experience.	2024-01-12 03:35:31.691624+08	2024-01-13 02:40:03.94791+08	https://i.ytimg.com/vi/gjQ0MzWH4Ss/hqdefault.jpg
10	Fluffy Panda Groove	An upbeat pop band from Beijing, their catchy tunes are inspired by panda antics, and they often perform in panda-themed outfits.	2024-01-12 03:35:55.444806+08	2024-01-13 02:42:18.479358+08	https://i.redd.it/qxcuvf8wuqc41.png
11	Bizarre Bagel Brigade	A jazz fusion ensemble from New York City, famous for their bagel-themed song titles and lyrics, often performing at local bagel shops.	2024-01-12 03:35:55.444806+08	2024-01-13 02:42:18.479358+08	https://i1.sndcdn.com/avatars-000350550545-fzf721-t500x500.jpg
12	Retro Robot Rumba	A synth-pop band from Berlin, embracing a retro-futuristic theme with robot-inspired dance moves and a flashy, metallic stage design.	2024-01-12 03:35:55.444806+08	2024-01-13 02:42:18.479358+08	https://cdn.dribbble.com/users/634508/screenshots/3697683/hideandseek_dribbble_2_still_2x.gif?resize=400x300&vertical=center
1	The Juggling Oranges	A vibrant indie-pop band hailing from a quaint Spanish town, known for incorporating actual orange juggling into their energetic live performances.	2024-01-08 14:48:24.469099+08	2024-01-13 02:37:34.3071+08	https://img.freepik.com/free-photo/smiling-woman-juggling-with-oranges_23-2148332164.jpg
2	Frankly Cookies	Frankly Cookies is a band that is based in Cookieland and specialises in the Cookie genre.	2024-01-08 16:01:14.549442+08	2024-01-13 02:37:34.3071+08	https://t4.ftcdn.net/jpg/02/63/38/89/360_F_263388995_LmTEjd9U2tPaFpE8s1NoxlaYk0kHqZn2.jpg
3	Midnight Taco Serenade	A lively Latin fusion band from the heart of Mexico City, known for their late-night performances in taco joints, blending mariachi influences with modern rhythms.	2024-01-12 18:37:03.525568+08	2024-01-13 02:37:34.3071+08	https://upload.wikimedia.org/wikipedia/commons/thumb/7/73/001_Tacos_de_carnitas%2C_carne_asada_y_al_pastor.jpg/1200px-001_Tacos_de_carnitas%2C_carne_asada_y_al_pastor.jpg
5	Neon Giraffe Parade	This psychedelic funk band, based in Amsterdam, is famous for their tall stilt performers dressed as neon giraffes who dance during concerts.	2024-01-12 03:35:09.2022+08	2024-01-13 02:38:55.775406+08	https://inst-1.cdn.shockers.de/hs_cdn/out/pictures/master/product/1/lustige-giraffe-kinderkostuem--baby-verkleidung-fasching--tierkostuem-giraffe-safari-baby--16816.jpg
6	Velvet Sofa Rebellion	Originating from a cozy coffee shop in Dublin, this folk-rock group is renowned for their songs about comfy furniture and relaxed living.	2024-01-12 03:35:31.691624+08	2024-01-13 02:38:55.775406+08	https://us.images.westend61.de/0000812362pw/portrait-of-group-of-female-friends-sitting-on-sofa-in-living-room-GIOF03422.jpg
9	Polka Dot Pirates	This trending sea shanty group from the coasts of Cornwall is known for their polka dot costumes and pirate-themed, accordion-heavy tunes. 	2024-01-12 03:35:31.691624+08	2024-01-21 18:20:16.087566+08	https://www.thoughtco.com/thmb/Z8KLz_N1h5ZJqS7iCSvYzfmJAhk=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/the-capture-of-the-pirate--blackbeard--1718-by-jean-leon-gerome-ferris-517200612-5a566ce1b39d0300378d4511.jpg
18	Neon Odessy	Neon Odyssey is a vibrant synth-pop band known for their retro-futuristic sound, blending 80s synthwave with modern electronic music. Their music is characterized by pulsating rhythms, neon-lit melodies, and nostalgic yet fresh vocals. 	2024-01-21 20:28:00.564032+08	2024-01-21 23:45:10.431569+08	https://store-images.s-microsoft.com/image/apps.4323.13568911399767367.40f48bf2-3e0d-4d5e-82ae-91d8c14bd585.f47ca390-9983-486a-834a-21f19b1ad04b?q=90&w=480&h=270
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, type, created_at, password_hash) FROM stdin;
a159490c-893c-4957-a318-7ad794c60e2b	yizhong187	normal	2024-01-16 16:39:32.579648+08	$2a$14$lvQvobBbBYOh9a7IaZ8cCOOwuoN5ztRlAsNSHIvRjx4XvVcBEIG1S
8f755116-0a7f-441b-b4cc-ee8eba8c11bc	god	super	2024-01-19 00:05:33.298153+08	$2a$14$Ce9GNM9jMpueV1UfZo7r4.6blghZvom2CO9HK4O8xI2HkUzAYdIKS
7e2f5b9c-763c-4519-b6fc-91f34f90eb48	timmy	normal	2024-01-22 11:45:54.465229+08	$2a$14$xdozi6K8DQ3WrV7FfEC13.iOLX3pUqteE8yiKbD0d5GKT3mtur8/q
\.


--
-- Data for Name: threads; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.threads (id, subforum_id, title, content, created_by, is_pinned, created_at, updated_at) FROM stdin;
106	11	I love the triple Bs!	They are a fantastic band!	a159490c-893c-4957-a318-7ad794c60e2b	f	2024-01-22 11:45:10.997976+08	2024-01-22 11:45:10.997976+08
107	11	New album review	Bizarre Bagel Brigade just released a new album! Reply your reviews below!	7e2f5b9c-763c-4519-b6fc-91f34f90eb48	f	2024-01-22 11:47:20.019799+08	2024-01-22 11:47:20.019799+08
\.


--
-- Data for Name: replies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.replies (id, thread_id, content, created_by, created_at, updated_at) FROM stdin;
83	107	I rate it a 5/5!	7e2f5b9c-763c-4519-b6fc-91f34f90eb48	2024-01-22 11:49:06.262248+08	2024-01-22 11:49:06.262248+08
84	107	I loved it as well!	a159490c-893c-4957-a318-7ad794c60e2b	2024-01-22 11:59:51.034769+08	2024-01-22 11:59:51.034769+08
\.


--
-- Name: replies_replyid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.replies_replyid_seq', 84, true);


--
-- Name: subforums_subforumid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.subforums_subforumid_seq', 20, true);


--
-- Name: threads_threadid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.threads_threadid_seq', 107, true);


--
-- PostgreSQL database dump complete
--

