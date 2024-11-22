CREATE TABLE "songs" (
  "id" uuid DEFAULT gen_random_uuid(),
  "group" varchar NOT NULL,
  "name" varchar NOT NULL,
  "release_date" varchar NOT NULL,
  "text" TEXT NOT NULL,
  "link" varchar NOT NULL, 
  PRIMARY KEY (id)
);

INSERT INTO songs ("group", "name", "release_date", "text", "link")
VALUES 
  ('The Beatles', 'Let It Be', '16.07.1970',
    E'When I find myself in times of trouble, Mother Mary comes to me\n\
Speaking words of wisdom, let it be\n\
And in my hour of darkness she is standing right in front of me\n\
Speaking words of wisdom, let it be\n\
Let it be, let it be, let it be, let it be\n\
Whisper words of wisdom, let it be\n',
    'https://www.imdb.com/name/nm1397313/?ref_=ls_t_1'),

  ('The Rolling Stones', 'Gimme Shelter', '12.08.1969',
    E'Ooh, a storm is threatening\n\
My very life today\n\
If I don\'t get some shelter\n\
Ooh yeah, I\'m gonna fade away\n\
War, children, it\'s just a shot away\n\
It\'s just a shot away\n\
War, children, it\'s just a shot away\n\
It\'s just a shot away\n',
    'https://www.imdb.com/name/nm1213869/?ref_=ls_t_2'),

  ('Led Zeppelin', 'Stairway to Heaven', '12.09.1971',
    E'There\'s a lady who\'s sure all that glitters is gold\n\
And she\'s buying a stairway to Heaven\n\
When she gets there she knows, if the stores are all closed\n\
With a word she can get what she came for\n\
Ooh, ooh, and she\'s buying a stairway to Heaven\n',
    'https://www.imdb.com/name/nm1213869/?ref_=ls_t_2'),

  ('The Doors', 'Riders on the Storm', '14.10.1971',
    E'Riders on the storm\n\
Riders on the storm\n\
Into this house, we\'re born\n\
Into this world, we\'re thrown\n\
Like a dog without a bone\n\
An actor out on loan\n\
Riders on the storm\n',
    'https://www.imdb.com/name/nm1213869/?ref_=ls_t_2'),

  ('Pink Floyd', 'Time', '20.11.1973',
    E'Ticking away the moments that make up a dull day\n\
You fritter and waste the hours in an off-hand way\n\
Kicking around on a piece of ground in your home town\n\
Waiting for someone or something to show you the way\n\
Tired of lying in the sunshine, staying home to watch the rain\n\
And you are young and life is long, and there is time to kill today\n\
And then one day you find ten years have got behind you\n\
No one told you when to run, you missed the starting gun\n',
    'https://www.imdb.com/name/nm1213869/?ref_=ls_t_2'),

  ('The Beach Boys', 'California Dreamin', '22.01.1986',
    E'All the leaves are brown\n\
And the sky is grey\n\
I\'ve been for a walk\n\
On a winter\'s day\n\
I\'d be safe and warm\n\
If I was in L.A\n',
    'https://www.imdb.com/name/nm1213869/?ref_=ls_t_2'),

  ('The Who', 'My Generation', '22.11.1965',
    E'People try to put us d-down (talkin\' \'bout my generation)\n\
Just because we get around (talkin\' \'bout my generation)\n\
Things they do look awful c-cold (talkin\' \'bout my generation)\n\
I hope I die before I get old (talkin\' \'bout my generation)\n\
This is my generation\n\
This is my generation, baby\n',
    'https://www.imdb.com/name/nm1213869/?ref_=ls_t_2'),

  ('U2', 'One', '24.12.1991',
    E'Is it getting better?\n\
Or do you feel the same?\n\
Will it make it easier on you now?\n\
You got someone to blame\n\
You say, one love, one life\n\
When it\'s one need in the night\n\
One love, we get to share it\n\
Leaves you baby if you don\'t care for it\n',
    'https://www.imdb.com/name/nm1213869/?ref_=ls_t_2'),

  ('Cream', 'White Room', '24.12.1991',
    E'In a white room with black curtains in the station\n\
Black roof country, no gold pavements, tired starlings\n\
Silver horses ran down moonbeams in your dark eyes\n\
Dawn light smiles on you leaving, my contentment\n\
I\'ll wait in this place where the sun never shines\n\
Wait in this place where the shadows run from themselves\n',
    'https://www.imdb.com/name/nm1213869/?ref_=ls_t_2'),

  ('The Velvet Underground', 'Pale Blue Eyes', '24.12.1969',
    E'Sometimes I feel so happy\n\
Sometimes I feel so sad\n\
Sometimes I feel so happy\n\
But mostly you just make me mad\n\
Baby, you just make me mad\n\
Linger on your pale blue eyes\n',
    'https://www.imdb.com/name/nm1213869/?ref_=ls_t_2'),

  ('The Kinks', 'Waterloo Sunset', '24.02.1967',
    E'Dirty old river, must you keep rolling\n\
Flowing into the night?\n\
People so busy, make me feel dizzy\n\
Taxi light shines so bright\n\
But I don\'t need no friends\n\
As long as I gaze on\n\
Waterloo sunset\n\
I am in paradise\n',
    'https://www.imdb.com/name/nm1213869/?ref_=ls_t_2'),

  ('Nirvana', 'Heart-Shaped Box', '24.12.1993',
    E'She eyes me like a Pisces when I am weak\n\
I\'ve been locked inside your heart-shaped box for weeks\n\
I\'ve been drawn into your magnet tar pit trap\n\
I wish I could eat your cancer when you turn black\n\
Hey\n\
Wait\n\
I got a new complaint\n',
    'https://www.imdb.com/name/nm1213869/?ref_=ls_t_2'),

  ('The Band Camino', 'Berenstein', '13.09.2007',
    E'You were always searching for deliverance\n\
Blatant misconceptions that you make\n\
My overall perception of this dream\n\
Is that I\'ll die before I wake\n',
    'https://www.imdb.com/name/nm1213869/?ref_=ls_t_2'),

  ('Crosby Stills Nash & Young', 'Ohio', '24.12.1971',
    E'Tin soldiers and Nixon coming\n\
We\'re finally on our own\n\
This summer I hear the drumming\n\
Four dead in Ohio\n\
Gotta get down to it, soldiers are cutting us down\n\
Should have been gone long ago\n\
What if you knew her and found her dead on the ground\n\
How can you run when you know?\n',
    'https://www.imdb.com/name/nm1213869/?ref_=ls_t_2');
