
<script>
    export let images;

    console.log(images)
</script>



<div class="flex w-full h-80 mt-4 overflow-x-auto rounded-lg whitespace-nowrap scroll-smooth transition delay-150
 ease-in-out">

    {#each images as image}

        <div class="flex">
            <img src="../{image}" alt="">
        </div>

        {/each}





</div>