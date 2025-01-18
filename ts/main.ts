import { CommandsClient } from './CommandsServiceClientPb';
import * as commands_pb from './commands_pb';

function output(message: string) {
  let element = <HTMLInputElement>document.getElementById('output');
  let timestamp = new Date().toISOString();
  element.innerText = timestamp + ': ' + message + '\n' + element.innerText;
}

function getGRPCSessionID(): string | null {
  let element = document.getElementById('grpc-session-id');
  return element ? element.innerText : null;
}

async function callCommands() {
  try {
    const metadata = {
      'session-id': getGRPCSessionID(),
    };
    let client = new CommandsClient('https://localhost:8443');

    // call unary rpc
    let unaryRequest = new commands_pb.UnaryRequest();
    unaryRequest.setMessage('Hi!');
    let unaryResponse = await client.unaryRPC(unaryRequest, metadata);
    output('Unary: ' + unaryResponse.getMessage());

    // call server stream rpc
    let serverStreamRequest = new commands_pb.ServerStreamRequest();
    serverStreamRequest.setMessage('Hello!');
    let serverStream = await client.serverStreamRPC(serverStreamRequest, metadata);
    serverStream.on('data', function (response: commands_pb.ServerStreamResponse) {
      output('Server Stream: ' + response.getMessage());
    });
  } catch (e) {
    output('Error: ' + e);
  }
}

const trigger = <HTMLInputElement>document.getElementById('trigger');
trigger.addEventListener('click', callCommands);