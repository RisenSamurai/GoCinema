<script>
    export let name = "Click me!";
    export let background = "bg-cinema-highlight";
    export let padding = "";
    export let width = "w-1/2";
    export let height = "h-10"

    if(padding === "") {
        padding = "p-4"
    }

</script>


<button on:click class="{background} {padding} {width} rounded-lg my-2 text-cinema-text font-bold">
    {name}
</button>