<script>
    export let title = "Movie";
    export let image = "";
</script>



<div class="flex flex-col">
    <div class="flex">
        <img src={image} alt={title} />
    </div>
    <span class="text-cinema-text">{title}</span>
</div>