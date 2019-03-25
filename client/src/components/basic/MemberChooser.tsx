import React, { useState } from "react";
import Toggle from "./Toggle";
import Button from "./Button";
import { Send, CID, Client } from "../../types";

type Props = {
  send: Send;
  me: CID;
  clients: Array<Client>;
  members: number;
};

function id<T>(x: T): T {
  return x;
}

export default ({ send, me, clients, members }: Props) => {
  const [checked, setChecked] = useState(() =>
    clients.map(({ CID }) => CID === me),
  );
  return (
    <div className="MemberChooser">
      {clients.map(({ CID, Name }, i) => (
        <Toggle
          key={CID}
          value={Name}
          checked={checked[i]}
          onChange={() => setChecked(checked.map((b, j) => (i === j ? !b : b)))}
        />
      ))}
      <Button
        type="submit"
        value="Submit"
        onClick={() => {
          const Members: Array<CID> = [];
          for (let i = 0; i < clients.length; i++) {
            if (checked[i]) {
              Members.push(clients[i].CID);
            }
          }
          send({ t: "MemberChoose", Members });
        }}
        disabled={checked.filter(id).length != members}
      />
    </div>
  );
};