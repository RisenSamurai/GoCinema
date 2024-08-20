<script>
    import Headers from "$lib/components/Headers.svelte";
    import MovieSlider from "$lib/components/MovieSlider.svelte";
    import Button from "$lib/components/Button.svelte";

    export let data;
    const { movie } = data;
    const { ratings } = data;
    console.log("Movie page!", movie);
    console.log("Ratings from TMDB!", ratings);



</script>



<section class="flex flex-col mt-2 px-2">

    <div class="flex flex-col bg-cinema-secondary rounded-lg mb-4">
        <div class="flex">
            <img class="object-cover rounded-lg w-full h-auto" src="../{movie.movie.poster}"
                 alt="{movie.movie.name}">
        </div>
    </div>

    <Headers title="{movie.movie.name}" extra={movie.movie.year}/>

    <MovieSlider images={data.movie.movie.images}/>

    <!--Ratings section -->

    <div class="flex flex-col bg-cinema-gold rounded-lg mt-4 pb-2">

        <span class="flex w-32 h-14 self-center"> <img src="/static/blue_short-tmdb.svg" alt="tmdb"> </span>
        <div class="flex w-full justify-around font-bold text-cinema-highlight">
            <span class="flex items-center">
                <svg xmlns="http://www.w3.org/2000/svg"
                                                 viewBox="0 0 24 24" fill="currentColor" class="size-5">
                    <path d="M7.493 18.5c-.425 0-.82-.236-.975-.632A7.48 7.48 0 0 1
   6 15.125c0-1.75.599-3.358 1.602-4.634.151-.192.373-.309.6-.397.473-.183.89-.514 1.212-.924a9.042
    9.042 0 0 1 2.861-2.4c.723-.384 1.35-.956 1.653-1.715a4.498 4.498 0 0 0 .322-1.672V2.75A.75.75 0 0
     1 15 2a2.25 2.25 0 0 1 2.25 2.25c0 1.152-.26 2.243-.723 3.218-.266.558.107 1.282.725 1.282h3.126c1.026
      0 1.945.694 2.054 1.715.045.422.068.85.068 1.285a11.95 11.95 0 0 1-2.649 7.521c-.388.482-.987
      .729-1.605.729H14.23c-.483 0-.964-.078-1.423-.23l-3.114-1.04a4.501 4.501 0 0 0-1.423-.23h-.777ZM2.
      331 10.727a11.969 11.969 0 0 0-.831 4.398 12 12 0 0 0 .52 3.507C2.28 19.482 3.105 20 3.994 20H4.9c.445
       0 .72-.498.523-.898a8.963 8.963 0 0 1-.924-3.977c0-1.708.476-3.305 1.302-4.666.245-.403-.028-.959-.5-.959H4
       .25c-.832 0-1.612.453-1.918 1.227Z" />
                </svg>
                {movie.ratings.vote_average}
            </span>
            <span class="flex items-center">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-6">
  <path fill-rule="evenodd" d="M15.22 6.268a.75.75 0 0 1 .968-.431l5.942 2.28a.75.75 0 0 1 .431.97l-2.28 5.94a.75.75 0 1 1-1.4-.537l1.63-4.251-1.086.484a11.2 11.2 0 0 0-5.45 5.173.75.75 0 0 1-1.199.19L9 12.312l-6.22 6.22a.75.75 0 0 1-1.06-1.061l6.75-6.75a.75.75 0 0 1 1.06 0l3.606 3.606a12.695 12.695 0 0 1 5.68-4.974l1.086-.483-4.251-1.632a.75.75 0 0 1-.432-.97Z" clip-rule="evenodd" />
</svg>
                {movie.ratings.popularity}
            </span>
            <span class="flex items-center">
                <svg xmlns="http://www.w3.org/2000/svg"
                                                 viewBox="0 0 24 24" fill="currentColor" class="size-5">
                    <path d="M4.5 6.375a4.125 4.125 0 1 1 8.25 0
   4.125 4.125 0 0 1-8.25 0ZM14.25 8.625a3.375 3.375 0 1 1 6.75 0 3.375 3.375 0 0 1-6.75 0ZM1.5 19.125a7.125 7.125 0 0 1 14.25 0v.003l-.001.119a.75.75 0 0 1-.363.63 13.067 13.067 0 0 1-6.761 1.873c-2.472 0-4.786-.684-6.76-1.873a.75.75 0 0 1-.364-.63l-.001-.122ZM17.25 19.128l-.001.144a2.25 2.25 0 0 1-.233.96 10.088 10.088 0 0 0 5.06-1.01.75.75 0 0 0 .42-.643 4.875 4.875 0 0 0-6.957-4.611 8.586 8.586 0 0 1 1.71 5.157v.003Z"/>
                </svg>
                {movie.ratings.vote_count}
            </span>
        </div>
    </div>


    <!--Description and general info section -->

    <div class="flex flex-col bg-cinema-secondary mt-4 rounded-lg">
        <div class="grid grid-cols-2 gap-2 font-semibold text-lg p-2 justify-center items-center">
            <div class="flex w-full mb-1">
                <span class="text-cinema-text text-right">Year:</span>
            </div>
            <div class="flex"><span class="text-cinema-text">{movie.movie.year}</span></div>
            <div class="flex w-full mb-1">
                <span class="text-cinema-text">Budget:</span>
            </div>
            <div>
                <span class="text-cinema-text">${movie.ratings.budget}</span>
            </div>
            <div class="flex w-full mb-1">
                <span class="text-cinema-text">Language: </span>
            </div>
            <div>
                <span class="text-cinema-text">{movie.movie.language}</span>
            </div>
            <div class="flex w-full mb-1">
                <span class="text-cinema-text">Release Date: </span>
            </div>
            <div>
                <span class="text-cinema-text">{movie.ratings.release_date}</span>
            </div>
            <div class="flex w-full mb-1">
                <span class="text-cinema-text">Runtime: </span>
            </div>
            <div>
                <span class="text-cinema-text">{movie.ratings.runtime}m</span>
            </div>
            <div class="flex  w-full mb-1">
                <span class="text-cinema-text">Directors: </span>
            </div>
            <div>
                <span class="text-cinema-text">{#each movie.movie.directors as director}
                    <span>{director}</span>
                {/each}</span>
            </div>
            <div class="flex  w-full mb-1">
                <span class="text-cinema-text">Writers: </span>
            </div>
            <div>
                <span class="text-cinema-text flex flex-col">{#each movie.movie.writers as writer}
                    <span class="flex">{writer}</span>
                {/each}</span>
            </div>
            <div class="flex  w-full mb-1">
                <span class="text-cinema-text">Producers: </span>
            </div>
            <div>
                <span class="text-cinema-text flex flex-col">{#each movie.movie.producers as producer}
                    <span class="flex">{producer}</span>
                {/each}</span>
            </div>
            <div class="flex w-full mb-1">
                <span class="text-cinema-text">Editors: </span>
            </div>
            <div>
                <span class="text-cinema-text flex flex-col">{#each movie.movie.editors as editor}
                    <span class="flex">{editor}</span>
                {/each}</span>
            </div>
            <div class="flex w-full mb-1">
                <span class="text-cinema-text">Cameras: </span>
            </div>
            <div>
                <span class="text-cinema-text flex flex-col">{#each movie.movie.cameras as camera}
                    <span class="flex">{camera}</span>
                {/each}</span>
            </div>
            <div class="flex w-full mb-1">
                <span class="text-cinema-text">Genres: </span>
            </div>
            <div>
                <span class="text-cinema-text flex flex-col">{#each movie.movie.genres as genre}
                    <a href={genre} class="flex">{genre}</a>
                {/each}</span>
            </div>
            <div class="flex w-full mb-1">
                <span class="text-cinema-text">Actors: </span>
            </div>
            <div>
                <span class="text-cinema-text flex flex-col">{#each movie.movie.actors as actor}
                    <a href={actor} class="flex">{actor}</a>
                {/each}</span>
            </div>
            <div class="flex w-full mb-1">
                <span class="text-cinema-text">Countries: </span>
            </div>
            <div>
                <span class="text-cinema-text flex flex-col">{#each movie.movie.countries as country}
                    <a href={country} class="flex">{country}</a>
                {/each}</span>
            </div>



        </div>
    </div>


    <div class="flex flex-col mt-2 bg-cinema-secondary rounded-lg">
        <Headers title="Description"/>

        <p class="text-cinema-text justify-center items-center
         hyphens-auto leading-normal break-words text-justify p-2">
            {@html movie.ratings.overview}
        </p>

    </div>

    <div class="flex flex-col mt-2 bg-cinema-secondary">
        <Headers title="Reviews"/>
        <div class="flex flex-col">
            <h3 class="text-cinema-text mt-2 font-bold">Leave a review</h3>
            <textarea class="rounded-lg"  rows="5"></textarea>
            <Button name="Send"  padding="p-2"/>
        </div>
        <div class="flex flex-col">
            <div class="flex flex-col">

                <div class="flex text-cinema-text items-center font-bold justify-between p-2 ">
                    <div class="flex justify-center h-16 w-16">
                        <img class="rounded-full object-cover w-full h-auto" src="/images/articles/9KKCxXsBPOpRXASuwRCF6ZUpfCy.webp" alt=""/>
                    </div>
                    <span>Tom Cruise</span>
                    <span class="italic">1/1/2024</span>
                </div>

                <div class="flex border-b border-b-cinema-highlight">
                    <p class="justify-center items-center text-cinema-text
                    hyphens-auto leading-normal break-words text-justify p-2">
                        Lorem ipsum dolor sit amet consectetur adipisicing elit.
                        Eum consectetur dignissimos fuga, provident at unde a possimus praesentium,
                        earum error doloribus sint deleniti id! Fugiat illo repellat cumque molestias voluptatem!
                    </p>
                </div>

                <div class="flex text-cinema-text items-center font-bold justify-between p-2">
                    <div class="flex justify-center h-16 w-16">
                        <img class="rounded-full object-cover w-full h-auto" src="/images/articles/9KKCxXsBPOpRXASuwRCF6ZUpfCy.webp" alt=""/>
                    </div>
                    <span>Tom Cruise</span>
                    <span class="italic">1/1/2024</span>
                </div>

                <div class="flex border-b border-b-cinema-highlight">
                    <p class="justify-center items-center text-cinema-text
                    hyphens-auto leading-normal break-words text-justify p-2">
                        Lorem ipsum dolor sit amet consectetur adipisicing elit.
                        Eum consectetur dignissimos fuga, provident at unde a possimus praesentium,
                        earum error doloribus sint deleniti id! Fugiat illo repellat cumque molestias voluptatem!
                    </p>
                </div>

            </div>
        </div>
    </div>

</section>