function addIngredient(name) {
    table = $('#ingredients_table');
    $('<li>', { class: 'ingredient-elem' })
        .append($('<span>', { class: 'name' }).html(name))
        .append($('<span>', { class: 'remove' }).html('[✗]'))
        .appendTo(table);
}

// Initializes form to submit new ingredient
function initForm() {
    $('#new_ingredient_form').submit((event) => {
        event.preventDefault();
        req_url = '/ingredient/new?' + $('#new_ingredient_form').serialize();
        console.log(req_url);
        post = $.post(req_url);

        post.fail(formRejected);
        post.done(formAccepted);
    });
};

// Action for failed post request
function formRejected(data) {
    console.log(data);
}

// Action for accepted post request
function formAccepted() {
    nameInput = $('#new_ingredient_form input[name=name]');
    addIngredient(nameInput.val());
    nameInput.val('');
    $('html').scrollTop($(document).height());
}

// Callbacks to delete ingredient
function initDeleteActions() {
    $('span.remove').each(function(index) {
        $(this).click(function() {
            id = $(this).parent().data('id');
            request = $.ajax({
                url: '/ingredient/' + id,
                type: 'DELETE'
            });
            request.fail(deleteRejected);
            request.done(deleteAccepetd(id));
        });
    });
}

function deleteRejected(data) {
    console.log(data);
}

function deleteAccepetd(id) {
    return function(data) {
        $('li.ingredient-elem[data-id="'+id+'"]').remove();
    };
}

// Entry point on page load
$( () => {
    initForm();
    initDeleteActions();
});
