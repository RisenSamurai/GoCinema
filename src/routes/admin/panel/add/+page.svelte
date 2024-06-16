
<script lang="ts">

import Headers from "$lib/components/Headers.svelte";
import Button from "$lib/components/Button.svelte";
import ActorForm from "$lib/components/ActorForm.svelte";

let cardType: string = "type";

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

function changeCardType(card: string) {
    cardType = card.toLowerCase();
}

</script>

<svelte:head>
    <title>GC | Add </title>
</svelte:head>


<section class="flex flex-col p-2 items-center">
    <Headers title="New Card" />

    <div class="flex flex-col justify-center w-full items-center mt-4">

        {#if cardType === "movie"}
            <div class="">
                <Headers title="Add Movie"/>
            </div>
            {:else if cardType === "series"}
                <div class="">
                    <Headers title="Add Series"/>
                </div>
            {:else if cardType === "article"}
                <div class="">
                    <Headers title="Add Article"/>
                </div>
            {:else if cardType === "actor"}
                <div class="flex w-full">
                    <ActorForm />
                </div>

        {/if}
        {#each buttons as button}
            <Button name="{button.name}" on:click={() => changeCardType(button.name)}/>

            {/each}



    </div>

</section>