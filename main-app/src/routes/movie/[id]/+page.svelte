<script>
    import Headers from "$lib/components/Headers.svelte";

    export let data;
    const { movie } = data;
    console.log("Movie page!", movie);
</script>



<section class="flex flex-col mt-6">
    <Headers title="{movie.name}"/>
</section>