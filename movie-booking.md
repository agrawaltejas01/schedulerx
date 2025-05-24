How would you design and implement a movie ticket booking service that manages

1. user registrations,
2. movie screenings with seat availability,
3. and ticket purchases?

You may choose any approach for data modeling,
handling seat availability constraints, and defining your API.
Please provide an end-to-end overview from data structures to workflow demonstrations showing how bookings are made and confirmed.

Model -
Theater
N screens
M Show, S seats

User

Bookings (S, M, U)

Numbers -

1. 10000 Theaters
2. 3 screens
3. 12 movies / day
4. 100 seats

10000 _ 3 _ 12 \* 100 = 36000000 ~ 1M

Theater

Flow -

Home screen -
JBTD

1. Explore shows
2. Filter Theater by location /add
   - Dates
   - Movies
3. Ezplore movies
   - Theaters along with time
4. Register
5. Trending movies

- /GET trendingMovies
- Movie is booked > 500, it should be trending
-

User clicks on search -
/GET search
