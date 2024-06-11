<script>

    import MovieCard from "$lib/components/MovieCard.svelte";

    export let header = "Header";
    export let background = "bg-cinema-secondary";

    let movies = [
        {
            title: "The Fall Guy",
            poster: ""
        }
    ]

</script>



<section class="flex flex-col mt-6">
    <h1 class="text-cinema-highlight text-2xl border-b border-b-cinema-silver pb-2 w-1/2">{header}</h1>

    <div class="flex {background} w-full h-64 mt-4 overflow-x-scroll rounded-lg">

        {#each }

        <MovieCard  />


    </div>

</section>