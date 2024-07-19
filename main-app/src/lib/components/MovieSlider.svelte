
<script>
    export let images;

    console.log(images)
</script>



<div class="flex w-full h-80 mt-4 overflow-x-auto rounded-lg whitespace-nowrap scroll-smooth transition delay-150
 ease-in-out">

    {#each images as image}

        <div class="flex flex-col shrink-0 w-1/2 rounded-lg p-2">
            <img class="object-cover rounded-lg w-full h-auto" src="../{image}" alt="">
        </div>

        {/each}





</div>