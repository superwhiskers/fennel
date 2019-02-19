// fennel.c - test code for libfennel
#include <stdio.h>
#include <argp.h>
#include "fennel.h"

const char *argp_program_version = "fennel-c v1";
const char *argp_program_bug_address = "<whiskerdev@protonmail.com>";
static char doc[] = "fennel-c - test program in c for the fennel c bindings";
static char args_doc[] = "nnid";
static struct argp_option options[] = {};

struct arguments {
  char *args[1];
};

static error_t parse_opt(int key, char *arg, struct argp_state *state) {

  struct arguments *arguments = state->input;

  switch (key) {
  case ARGP_KEY_ARG:
    if (state->arg_num >= 1) {

      argp_usage(state);

    }

    arguments->args[state->arg_num] = arg;
    break;

  case ARGP_KEY_END:
    if (state->arg_num < 1) {

      argp_usage(state);

    }
    break;

  default:
      return ARGP_ERR_UNKNOWN;

  }
  return 0;

}

static struct argp argp = {options, parse_opt, args_doc, doc};

int main(int argc, char **argv) {

  struct arguments arguments;

  argp_parse(&argp, argc, argv, 0, 0, &arguments);
  
  struct fennel_ClientInformation clientInfo;
  clientInfo.ClientID = "ea25c66c26b403376b4c5ed94ab9cdea";
  clientInfo.ClientSecret = "d137be62cb6a2b831cad8c013b92fb55";
  clientInfo.DeviceCert = "";
  clientInfo.Environment = "";
  clientInfo.Country = "US";
  clientInfo.Region = "2";
  clientInfo.SysVersion = "1111";
  clientInfo.Serial = "1";
  clientInfo.DeviceID = "1";
  clientInfo.DeviceType = "";
  clientInfo.PlatformID = "1";

  struct fennel_Error *error = malloc(sizeof error);
  fennel_AccountServerClient client;
  client = fennel_newAccountServerClient("https://account.nintendo.net/v1/api", "keypair/ctr-common-cert.pem", "keypair/ctr-common-key.pem", clientInfo, &error);

  if ((*error).Type == fennel_ErrorTypeError) {

    printf("error: %s\n", (*error).Error);
    free(error);
    return 1;

  }
  free(error);
  
  if (argc < 2) {

    printf("no parameter was provided\n");
    return 1;

  }

  error = malloc(sizeof error);
  int exists = fennel_accountServerClient_doesUserExist(
      client, arguments.args[0], &error);
  if ((*error).Type == fennel_ErrorTypeError) {

    printf("error: %s\n", (*error).Error);
    free(error);
    return 1;

  } else if ((*error).Type == fennel_ErrorTypeErrorXML) {

    printf("error: %s\n", (*error).ErrorXML.Message);
    free(error);
    return 1;

  }
  free(error);
  
  if (exists == 0) {

    printf("no, the user does not exist\n");

  } else {

    printf("yes, the user does exist\n");

  }
  
}
