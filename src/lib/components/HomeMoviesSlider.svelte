<script>

    import MovieCard from "$lib/components/MovieCard.svelte";
    import NewsCard from "$lib/components/NewsCard.svelte";

    export let header = "Header";
    export let background = "bg-cinema-secondary";
    export let type = "movies";
    export let items;




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

    let articles = [
        {
            title: "Tom Cruise and his new movie",
            poster: "",
            content: "Tom Cruise is going to release his new movie in the next year.",
        }
    ]

    console.log(items)

</script>



<section class="flex flex-col mt-6">
    <h1 class="text-cinema-highlight text-2xl border-b border-b-cinema-silver pb-2 w-1/2">{header}</h1>

    <div class="flex {background} w-full h-80 mt-4 overflow-x-auto rounded-lg whitespace-nowrap scroll-smooth transition delay-150 ease-in-out">

        {#if type == "movies"}
                {#each items as item}
                    <MovieCard  title="{item.name}" poster="{item.poster}" link={"/movie/"+item.id}/>
                {/each}
        {:else if type == "series"}
            {#each series as serial}
                <MovieCard  title="{serial.title}" poster="{serial.poster}"/>
            {/each}
            {:else}
            {#each articles as article}

                <NewsCard  title="{article.title}" poster="{article.poster}" content="{article.content}"/>

            {/each}

                {/if}



    </div>

</section>