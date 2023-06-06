CREATE TABLE public.users (
      id bigserial NOT NULL,
      created_at timestamptz NULL,
      updated_at timestamptz NULL,
      username text NOT NULL,
      email text NOT NULL,
      "password" text NOT NULL,
      age int8 NOT NULL,
      CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE UNIQUE INDEX idx_users_email ON public.users USING btree (email);
CREATE UNIQUE INDEX idx_users_username ON public.users USING btree (username);

CREATE TABLE public.photos (
       id bigserial NOT NULL,
       created_at timestamptz NULL,
       updated_at timestamptz NULL,
       title text NOT NULL,
       caption text NULL,
       photo_url text NOT NULL,
       user_id int8 NULL,
       CONSTRAINT photos_pkey PRIMARY KEY (id),
       CONSTRAINT fk_photos_user FOREIGN KEY (user_id) REFERENCES public.users(id)
);

CREATE TABLE public."comments" (
       id bigserial NOT NULL,
       created_at timestamptz NULL,
       updated_at timestamptz NULL,
       user_id int8 NULL,
       photo_id int8 NULL,
       message text NOT NULL,
       CONSTRAINT comments_pkey PRIMARY KEY (id),
       CONSTRAINT fk_comments_photo FOREIGN KEY (photo_id) REFERENCES public.photos(id),
       CONSTRAINT fk_comments_user FOREIGN KEY (user_id) REFERENCES public.users(id)
);

CREATE TABLE public.social_media (
     id bigserial NOT NULL,
     created_at timestamptz NULL,
     updated_at timestamptz NULL,
     "name" text NOT NULL,
     social_media_url text NOT NULL,
     user_id int8 NULL,
     CONSTRAINT social_media_pkey PRIMARY KEY (id),
     CONSTRAINT fk_social_media_user FOREIGN KEY (user_id) REFERENCES public.users(id)
);