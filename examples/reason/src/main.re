/* server_example.ml */
open Lwt;
open Cohttp;
open Cohttp_lwt_unix;

let server = {
  let callback = (_conn, req, body) => {
    let uri = req |> Request.uri |> Uri.to_string;
    let meth = req |> Request.meth |> Code.string_of_method;
    let headers = req |> Request.headers |> Header.to_string;
    body
    |> Cohttp_lwt.Body.to_string
    >|= (
      body =>
        Printf.sprintf(
          "Uri: %s\nMethod: %s\nHeaders\nHeaders: %s\nBody: %s",
          uri,
          meth,
          headers,
          body,
        )
    )
    >>= (body => Server.respond_string(~status=`OK, ~body, ()));
  };

  Server.create(~mode=`TCP(`Port(8000)), Server.make(~callback, ()));
};

let () = ignore(Lwt_main.run(server));
