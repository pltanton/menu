function addIngredient(name) {
    table = $('#ingredients-table');
    $('<tr>').append($('<td>').html(name)).appendTo(table)
}

function initForm() {
    $('#new-ingredient-form').submit((event) => {
        event.preventDefault();
        req_url = '/ingredient/new?' + $('#new-ingredient-form').serialize();
        console.log(req_url)
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
    nameInput = $('#new-ingredient-form input[name=name]');
    addIngredient(nameInput.val());
    nameInput.val('');
    $('html').scrollTop($(document).height());
}

$( () => {
    initForm();
});
