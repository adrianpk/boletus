# Changelog


## <a name="2019121"></a>20191221 - User signup confirmation processing


## <a name="20191219"></a>20191219 - Mailer implemented

Amazon SES mailer.

Confirmation email can be disabled by configuration but still email can be generated for debugging purposes.

## <a name="20191218"></a>20191218 - User signup

Confirmation email and verification link processing still pending.

## <a name="20191217"></a>20191217 - User signin

[User signin](screenshots.md#user-signin)

## <a name="20191216"></a>20191216 - Model validation and flash messages

Model validations moved to service layer.

[Model validation](screenshots.md#model-validation)

[Flash messages](screenshots.md#flash-messages)

## <a name="20191215"></a>20191215 - Form validations

Final intention is to move validations to the service layer or to another one that located at the "same depth".
In that way they could be reused on JSON REST and gRPC endpoints.

## <a name="20191213"></a>20191213 - RESTful actions completed

RESTful actions completed.

## <a name="20191128"></a>20191128 - Page text localization

Service web pages can now present a localized version of their texts depending on the regional configuration of the browser that accesses the service.

The service use the same translation files under '/assets/web/embed/i18n' to create language bundles.

## <a name="20191127"></a>20191127 - Simplified error handling

Common and repetitive operations on handler errors are now generalized in a single one.

## <a name="20191126"></a>20191126 - Simplified flash messages handling

Wrapped response now takes care to append new messages from current action to pending ones stored by other actions before a redirect.

## <a name="20191124"></a>20191124 - Embedded translations and form data session store

Localization files now are embedded in executable as any other resource.

Endpoint helpers to store, retrieve and clear form data accross requests.

Mainly used to avoid having to fill forms again in case of submission errors.

## <a name="20191123"></a>20191123 - Internationalization

Not fully implemented but basically this is how it works.

The I18N middleware currently tries to read the `lang` field submitted in POST, PUT, PATCH request; alternatively that same field value but read from the query string is used for all methods.

If still not found the value to be used to decide the language to be use will be associated to the one reported by `Accept-Language` request HTTP header.

In any other case the default language (English) will be used.

Soon, priority to decide the language will be the lang user chooses at sign-up time and/or another set by user using a gui switch.

A general clean-up is still needed, eventually a large part of the implementation will be moved to web module.

Finally, as with other assets, localization files under `assets/web/embed/i18n` will be embedded directly into the executable. Temporarily at this stage of development they are accessed from the filesytem.
