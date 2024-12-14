<script>
    export let data;
</script>

<svelte:head>
    <title>GoCinema | Article</title>
</svelte:head>



<section class="flex flex-col mt-2 px-2">

    <h1 class="text-2xl font-bold">{data.title}</h1>
    <div class="flex">
        <img src={data.image} alt="" class="w-16 h-16">
    </div>
    <div class="text-sm text-gray-500">{data.date}</div>
    <span class="text-sm text-gray-500">{data.author}</span>

    <div class="mt-2">{data.content}</div>

</section>


