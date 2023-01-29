--
-- PostgreSQL database dump
--

-- Dumped from database version 13.2
-- Dumped by pg_dump version 14.6

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
-- Name: timescaledb; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS timescaledb WITH SCHEMA public;


--
-- Name: EXTENSION timescaledb; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION timescaledb IS 'Enables scalable inserts and complex queries for time-series data';


--
-- Name: notify_tweet(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.notify_tweet() RETURNS trigger
    LANGUAGE plpgsql
    AS $$

	DECLARE
		data json;

	BEGIN
		
		-- create row to json
		data = row_to_json(NEW);

		-- notify tweets stream
		PERFORM pg_notify('tweet_stream', data::text);

		RETURN NULL;
	END;

$$;


ALTER FUNCTION public.notify_tweet() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: birds; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.birds (
    id integer NOT NULL,
    name text NOT NULL,
    entity_type integer NOT NULL,
    bird_fk integer NOT NULL,
    score double precision NOT NULL,
    n_positive integer NOT NULL,
    n_negative integer NOT NULL,
    img_url text NOT NULL
);


ALTER TABLE public.birds OWNER TO postgres;

--
-- Name: birds_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.birds_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.birds_id_seq OWNER TO postgres;

--
-- Name: birds_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.birds_id_seq OWNED BY public.birds.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO postgres;

--
-- Name: tweets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tweets (
    id integer NOT NULL,
    body text NOT NULL,
    bird_fk integer NOT NULL,
    author_name text NOT NULL,
    author_username text NOT NULL,
    post_time timestamp with time zone NOT NULL,
    score double precision NOT NULL,
    certainty double precision NOT NULL
);


ALTER TABLE public.tweets OWNER TO postgres;

--
-- Name: tweets_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tweets_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tweets_id_seq OWNER TO postgres;

--
-- Name: tweets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tweets_id_seq OWNED BY public.tweets.id;


--
-- Name: birds id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.birds ALTER COLUMN id SET DEFAULT nextval('public.birds_id_seq'::regclass);


--
-- Name: tweets id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tweets ALTER COLUMN id SET DEFAULT nextval('public.tweets_id_seq'::regclass);


--
-- Data for Name: cache_inval_bgw_job; Type: TABLE DATA; Schema: _timescaledb_cache; Owner: postgres
--

COPY _timescaledb_cache.cache_inval_bgw_job  FROM stdin;
\.


--
-- Data for Name: cache_inval_extension; Type: TABLE DATA; Schema: _timescaledb_cache; Owner: postgres
--

COPY _timescaledb_cache.cache_inval_extension  FROM stdin;
\.


--
-- Data for Name: cache_inval_hypertable; Type: TABLE DATA; Schema: _timescaledb_cache; Owner: postgres
--

COPY _timescaledb_cache.cache_inval_hypertable  FROM stdin;
\.


--
-- Data for Name: hypertable; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

COPY _timescaledb_catalog.hypertable (id, schema_name, table_name, associated_schema_name, associated_table_prefix, num_dimensions, chunk_sizing_func_schema, chunk_sizing_func_name, chunk_target_size, compression_state, compressed_hypertable_id, replication_factor) FROM stdin;
\.


--
-- Data for Name: chunk; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

COPY _timescaledb_catalog.chunk (id, hypertable_id, schema_name, table_name, compressed_chunk_id, dropped) FROM stdin;
\.


--
-- Data for Name: dimension; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

COPY _timescaledb_catalog.dimension (id, hypertable_id, column_name, column_type, aligned, num_slices, partitioning_func_schema, partitioning_func, interval_length, integer_now_func_schema, integer_now_func) FROM stdin;
\.


--
-- Data for Name: dimension_slice; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

COPY _timescaledb_catalog.dimension_slice (id, dimension_id, range_start, range_end) FROM stdin;
\.


--
-- Data for Name: chunk_constraint; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

COPY _timescaledb_catalog.chunk_constraint (chunk_id, dimension_slice_id, constraint_name, hypertable_constraint_name) FROM stdin;
\.


--
-- Data for Name: chunk_data_node; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

COPY _timescaledb_catalog.chunk_data_node (chunk_id, node_chunk_id, node_name) FROM stdin;
\.


--
-- Data for Name: chunk_index; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

COPY _timescaledb_catalog.chunk_index (chunk_id, index_name, hypertable_id, hypertable_index_name) FROM stdin;
\.


--
-- Data for Name: compression_chunk_size; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

COPY _timescaledb_catalog.compression_chunk_size (chunk_id, compressed_chunk_id, uncompressed_heap_size, uncompressed_toast_size, uncompressed_index_size, compressed_heap_size, compressed_toast_size, compressed_index_size, numrows_pre_compression, numrows_post_compression) FROM stdin;
\.


--
-- Data for Name: continuous_agg; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

COPY _timescaledb_catalog.continuous_agg (mat_hypertable_id, raw_hypertable_id, user_view_schema, user_view_name, partial_view_schema, partial_view_name, bucket_width, direct_view_schema, direct_view_name, materialized_only) FROM stdin;
\.


--
-- Data for Name: continuous_aggs_hypertable_invalidation_log; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

COPY _timescaledb_catalog.continuous_aggs_hypertable_invalidation_log (hypertable_id, lowest_modified_value, greatest_modified_value) FROM stdin;
\.


--
-- Data for Name: continuous_aggs_invalidation_threshold; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

COPY _timescaledb_catalog.continuous_aggs_invalidation_threshold (hypertable_id, watermark) FROM stdin;
\.


--
-- Data for Name: continuous_aggs_materialization_invalidation_log; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

COPY _timescaledb_catalog.continuous_aggs_materialization_invalidation_log (materialization_id, lowest_modified_value, greatest_modified_value) FROM stdin;
\.


--
-- Data for Name: hypertable_compression; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

COPY _timescaledb_catalog.hypertable_compression (hypertable_id, attname, compression_algorithm_id, segmentby_column_index, orderby_column_index, orderby_asc, orderby_nullsfirst) FROM stdin;
\.


--
-- Data for Name: hypertable_data_node; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

COPY _timescaledb_catalog.hypertable_data_node (hypertable_id, node_hypertable_id, node_name, block_chunks) FROM stdin;
\.


--
-- Data for Name: metadata; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

COPY _timescaledb_catalog.metadata (key, value, include_in_telemetry) FROM stdin;
exported_uuid	d27a32bc-0853-4c9a-a532-d4d6ef563177	t
\.


--
-- Data for Name: remote_txn; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

COPY _timescaledb_catalog.remote_txn (data_node_name, remote_transaction_id) FROM stdin;
\.


--
-- Data for Name: tablespace; Type: TABLE DATA; Schema: _timescaledb_catalog; Owner: postgres
--

COPY _timescaledb_catalog.tablespace (id, hypertable_id, tablespace_name) FROM stdin;
\.


--
-- Data for Name: bgw_job; Type: TABLE DATA; Schema: _timescaledb_config; Owner: postgres
--

COPY _timescaledb_config.bgw_job (id, application_name, schedule_interval, max_runtime, max_retries, retry_period, proc_schema, proc_name, owner, scheduled, hypertable_id, config) FROM stdin;
\.


--
-- Data for Name: birds; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.birds (id, name, entity_type, bird_fk, score, n_positive, n_negative, img_url) FROM stdin;
1	Elon Musk	0	1	0	0	0	https://upload.wikimedia.org/wikipedia/commons/3/34/Elon_Musk_Royal_Society_%28crop2%29.jpg
2	Sam Bankman-Fried	0	2	0	0	0	https://en.wikipedia.org/wiki/Sam_Bankman-Fried#/media/File:Sam_Bankman-Fried_2022.png
3	Kanye West	0	3	0	0	0	https://upload.wikimedia.org/wikipedia/commons/thumb/5/5c/Kanye_West_at_the_2009_Tribeca_Film_Festival_%28crop_2%29.jpg/800px-Kanye_West_at_the_2009_Tribeca_Film_Festival_%28crop_2%29.jpg
4	Andrew Tate	0	4	0	0	0	https://cdn.vox-cdn.com/thumbor/RcAdlMhw-adDQnJiVWKRPUSP10M=/0x0:2024x1038/1200x800/filters:focal(989x320:1311x642)/cdn.vox-cdn.com/uploads/chorus_image/image/71278865/Screen_Shot_2022_08_23_at_4.22.21_PM.0.png
5	Mark Zuckerberg	0	5	0	0	0	http://t2.gstatic.com/licensed-image?q=tbn:ANd9GcRYT8u3wI9a6gaceV_KUz_4LyVnTep01eLQa3xBOSBl-uBjA54InsDoUIZBiim-VGpQv2ad2OTZQ5hDa6U
6	Sam Altman	0	1	0	0	0	https://pbs.twimg.com/profile_images/804990434455887872/BG0Xh7Oa_400x400.jpg
7	Jeff Bezos	0	1	0	0	0	http://t3.gstatic.com/licensed-image?q=tbn:ANd9GcQcKtPg4LQ1A7_j_7_ph7FfTTTjQrnqOdC2EPUHdeqAZ01JOImw19i9gvYHROXo0HahI13E_dZ1ZekfGEE
8	Bill Gates	0	1	0	0	0	https://pbs.twimg.com/profile_images/1564398871996174336/M-hffw5a_400x400.jpg
9	Joe Rogan	0	1	0	0	0	https://pbs.twimg.com/profile_images/552307347851210752/vrXDcTFC_400x400.jpeg
10	Jack Dorsey	0	1	0	0	0	https://pbs.twimg.com/profile_images/1115644092329758721/AFjOr-K8_400x400.jpg
11	Tim Cook	0	1	0	0	0	https://pbs.twimg.com/profile_images/1535420431766671360/Pwq-1eJc_400x400.jpg
12	Lisa Su	0	1	0	0	0	https://pbs.twimg.com/profile_images/836654735553777664/gWyrS-2a_400x400.jpg
13	Sundar Pichai	0	1	0	0	0	https://pbs.twimg.com/profile_images/864282616597405701/M-FEJMZ0_400x400.jpg%27
14	Tesla	1	1	0	0	0	https://pbs.twimg.com/profile_images/1337607516008501250/6Ggc4S5n_400x400.png
15	FTX	1	1	0	0	0	https://pbs.twimg.com/profile_images/1587616813185392640/62LYZEro_400x400.jpg
16	Meta	1	1	0	0	0	https://pbs.twimg.com/profile_images/1453818753880190978/HqrrEcrI_400x400.png%27
17	Amazon	1	1	0	0	0	https://pbs.twimg.com/profile_images/1587352380454338561/Uav_0HHn_400x400.jpg
18	Apple	1	1	0	0	0	https://pbs.twimg.com/profile_images/1283958620359516160/p7zz5dxZ_400x400.jpg
19	AMD	1	1	0	0	0	https://pbs.twimg.com/profile_images/1534988679076368384/In8IVKDB_400x400.png
20	Google	1	1	0	0	0	https://pbs.twimg.com/profile_images/1453818753880190978/HqrrEcrI_400x400.png%27
21	Microsoft	1	1	0	0	0	https://pbs.twimg.com/profile_images/1585617858184323072/Vy138ToA_400x400.png
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schema_migrations (version, dirty) FROM stdin;
1	f
\.


--
-- Data for Name: tweets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tweets (id, body, bird_fk, author_name, author_username, post_time, score, certainty) FROM stdin;
\.


--
-- Name: chunk_constraint_name; Type: SEQUENCE SET; Schema: _timescaledb_catalog; Owner: postgres
--

SELECT pg_catalog.setval('_timescaledb_catalog.chunk_constraint_name', 1, false);


--
-- Name: chunk_id_seq; Type: SEQUENCE SET; Schema: _timescaledb_catalog; Owner: postgres
--

SELECT pg_catalog.setval('_timescaledb_catalog.chunk_id_seq', 1, false);


--
-- Name: dimension_id_seq; Type: SEQUENCE SET; Schema: _timescaledb_catalog; Owner: postgres
--

SELECT pg_catalog.setval('_timescaledb_catalog.dimension_id_seq', 1, false);


--
-- Name: dimension_slice_id_seq; Type: SEQUENCE SET; Schema: _timescaledb_catalog; Owner: postgres
--

SELECT pg_catalog.setval('_timescaledb_catalog.dimension_slice_id_seq', 1, false);


--
-- Name: hypertable_id_seq; Type: SEQUENCE SET; Schema: _timescaledb_catalog; Owner: postgres
--

SELECT pg_catalog.setval('_timescaledb_catalog.hypertable_id_seq', 1, false);


--
-- Name: bgw_job_id_seq; Type: SEQUENCE SET; Schema: _timescaledb_config; Owner: postgres
--

SELECT pg_catalog.setval('_timescaledb_config.bgw_job_id_seq', 1000, false);


--
-- Name: birds_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.birds_id_seq', 21, true);


--
-- Name: tweets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tweets_id_seq', 1, false);


--
-- Name: birds birds_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.birds
    ADD CONSTRAINT birds_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: tweets tweets_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tweets
    ADD CONSTRAINT tweets_pkey PRIMARY KEY (id);


--
-- Name: birds tweets_notify_event; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER tweets_notify_event AFTER UPDATE ON public.birds FOR EACH ROW EXECUTE FUNCTION public.notify_tweet();


--
-- Name: birds birds_bird_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.birds
    ADD CONSTRAINT birds_bird_fk_fkey FOREIGN KEY (bird_fk) REFERENCES public.birds(id);


--
-- Name: tweets tweets_bird_fk_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tweets
    ADD CONSTRAINT tweets_bird_fk_fkey FOREIGN KEY (bird_fk) REFERENCES public.birds(id);


--
-- PostgreSQL database dump complete
--

