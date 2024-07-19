<script>
    import Headers from "$lib/components/Headers.svelte";
    import MovieSlider from "$lib/components/MovieSlider.svelte";

    export let data;
    const { movie } = data;
    console.log("Movie page!", movie);
</script>



<section class="flex flex-col">

    <div class="flex flex-col bg-cinema-secondary">
        <div class="flex">
            <img class="object-cover rounded-lg w-full h-auto" src="../{movie.poster}" alt="{movie.name}">
        </div>
    </div>

    <Headers title="{movie.name}" extra={movie.year}/>

    <MovieSlider images={data.movie.images}/>

</section>