<script>
    import Headers from "$lib/components/Headers.svelte";
    import MovieSlider from "$lib/components/MovieSlider.svelte";

    export let data;
    const { movie } = data;
    const { ratings } = data;
    console.log("Movie page!", movie);
    console.log("Ratings from TMDB!", ratings);



</script>



<section class="flex flex-col mt-2 px-2">

    <div class="flex flex-col bg-cinema-secondary rounded-lg mb-4">
        <div class="flex">
            <img class="object-cover rounded-lg w-full h-auto" src="../{movie.poster}" alt="{movie.name}">
        </div>
    </div>

    <Headers title="{movie.name}" extra={movie.year}/>

    <MovieSlider images={data.movie.images}/>

    <!--Ratings section -->

    <div class="flex bg-cinema-gold rounded-lg p-4 mt-4">
        <span>{ratings.vote_average}</span>
        <span>Rating 2</span>
    </div>


    <!--Description and general info section -->

    <div class="flex flex-col bg-cinema-secondary mt-4 rounded-lg">
        <div class="flex flex-col font-semibold text-lg justify-between p-2">
            <div class="flex justify-between w-full mb-1">
                <span class="text-cinema-text">Year: </span>
                <span class="text-cinema-text">{movie.year}</span>
            </div>
            <div class="flex justify-between w-full mb-1">
                <span class="text-cinema-text">Budget: </span>
                <span class="text-cinema-text">${movie.budget}</span>
            </div>
            <div class="flex justify-between w-full mb-1">
                <span class="text-cinema-text">Language: </span>
                <span class="text-cinema-text">{movie.language}</span>
            </div>
            <div class="flex justify-between w-full mb-1">
                <span class="text-cinema-text">Release Date: </span>
                <span class="text-cinema-text">{movie.release_date}</span>
            </div>

        </div>
    </div>
</section>