
<script>
    import Headers from "$lib/components/Headers.svelte";

    let menu = [
        {
            title: "New Card",
            href: "panel/add/"
        },
        {
            title: "Edit Card",
            href: "panel/edit/"
        },
        {
            title: "Delete Card",
            href: "panel/delete/"
        }
    ]
</script>

<svelte:head>
    <title>GC | Panel</title>
</svelte:head>


<section class="flex flex-col p-2 items-center">
    <Headers title="CRUD"/>

    {#each menu as item}

        <div class="flex flex-col items-center bg-cinema-secondary w-full mt-4">
            <a class="flex bg-cinema-secondary text-lh font-bold text-cinema-text items-center justify-center p-2 w-full"
               href={item.href}>
                {item.title}
            </a>
        </div>


        {/each}



</section>