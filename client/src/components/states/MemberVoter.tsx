import React from "react";
import { Send, CID, Client } from "../../types";
import Voter from "../Voter";

type Props = {
  send: Send;
  me: CID;
  captain: CID;
  clients: Array<Client>;
  members: Array<CID>;
};

const modifiers = (cid: CID, me: CID, captain: CID): string =>
  cid === me && cid === captain
    ? " (you, captain)"
    : cid === me
    ? " (you)"
    : cid === captain
    ? " (captain)"
    : "";

const options: Array<[string, boolean]> = [
  ["Occur", true],
  ["Not occur", false],
];

export default ({ send, captain, me, clients, members }: Props) => (
  <div className="MemberVoter">
    <h1>Member vote</h1>
    <p>
      The captain, {clients.find(({ CID }) => CID === captain)!.Name}, has
      selected the following players to participate in the mission.
    </p>
    {clients
      .filter(({ CID }) => members.includes(CID))
      .map(({ CID, Name }) => (
        <div key={CID}>
          {Name}
          {modifiers(CID, me, captain)}
        </div>
      ))}
    <Voter
      prompt="Should the mission occur?"
      options={options}
      onVote={Vote => send({ t: "MemberVote", Vote })}
    />
  </div>
);
