
<script>

import Headers from "$lib/components/Headers.svelte";
import Button from "$lib/components/Button.svelte";
import ActorForm from "$lib/components/ActorForm.svelte";
import MovieForm from "$lib/components/MovieForm.svelte";
import ArticleForm from "$lib/components/ArticleForm.svelte";


let cardType= "type";
let menuOpen = true;


let buttons = [
    {
        name: "Movie",
    },
    {
        name: "Series",
    },
    {
        name: "Article",
    },
    {
        name: "Actor",
    }
];


function changeCardType(card, tumb) {
    cardType = card.toLowerCase();
    toggleMenu(tumb, card.toLowerCase())
}

function toggleMenu(tumb, card) {

    menuOpen = tumb;
    cardType = card;
}

</script>

<svelte:head>
    <title>GC | Add </title>
</svelte:head>


<section class="flex flex-col p-2 items-center h-screen">
    <Headers title="New Card" />

    <div class="flex flex-col justify-center w-full items-center mt-4">

        {#if cardType === "movie"}
            <div class="flex w-full flex-col">
                <Headers title="Add Movie"/>
                <MovieForm />
            </div>
            {:else if cardType === "series"}
                <div class="">
                    <Headers title="Add Series"/>
                </div>
            {:else if cardType === "article"}
                <div class=" flex flex-col w-full">
                    <Headers title="Add Article"/>
                    <ArticleForm />
                </div>
            {:else if cardType === "actor"}
                <div class="flex w-full flex-col">
                    <ActorForm />
                </div>

        {/if}

        {#if !menuOpen}
            <Button name={"Back"} on:click={() => toggleMenu(true, "type")} />
        {/if}

        {#if menuOpen}
            {#each buttons as button}
                <Button name="{button.name}" on:click={() => changeCardType(button.name, false)}/>

            {/each}
            {/if}




    </div>

</section>