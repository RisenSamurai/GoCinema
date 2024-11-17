<script>

    import MovieCard from "$lib/components/MovieCard.svelte";
    import NewsCard from "$lib/components/NewsCard.svelte";

    export let header = "Header";
    export let background = "bg-cinema-secondary";
    export let type = "movie";
    export let items;


    let baseImgUrl = "https://image.tmdb.org/t/p/";
    let imageSize = "w780";

    console.log(items);


    let series = [
        {
            title: "Supernatural",
            poster: "/images/posters/Supernatural/1fvjtNHks9SufUVv8XvhGXVpntj.webp",
        },
        {
            title: "Supernatural",
            poster: "/images/posters/Supernatural/1fvjtNHks9SufUVv8XvhGXVpntj.webp",
        },
        {
            title: "Supernatural",
            poster: "/images/posters/Supernatural/1fvjtNHks9SufUVv8XvhGXVpntj.webp",
        }
    ]




</script>



<section class="flex flex-col mt-6">
    <h1 class="text-cinema-highlight text-2xl border-b border-b-cinema-silver pb-2 w-1/2">{header}</h1>

    <div class="flex {background} w-full h-80 mt-4 overflow-x-auto rounded-lg whitespace-nowrap scroll-smooth transition delay-150 ease-in-out">

        {#if type === "movie"}
                {#each items as item}
                    <MovieCard  title="{item.title}" poster="{baseImgUrl}{imageSize}{item.poster_path}" link={"/movie/"+item.id}/>
                {/each}
        {:else if type === "series"}
            {#each series as serial}
                <MovieCard  title="{serial.title}" poster="{serial.poster}"/>
            {/each}
            {:else}
            {#each items as article}

                <NewsCard  title="{article.Title}" poster="{article.Image}" content="{article.Content}" id="{article.Id}"/>

            {/each}

                {/if}



    </div>

</section>