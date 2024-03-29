import React, { useState } from "react";
import { CID, Client } from "../../shared";
import { Translation, S } from "../../etc";
import Toggle from "./Toggle";
import Button from "./Button";

type Props = {
  t: Translation;
  send: S;
  me: CID;
  clients: Array<Client>;
  members: number;
};

function id<T>(x: T): T {
  return x;
}

export default ({ t, send, me, clients, members }: Props) => {
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
        value={t.submit}
        onClick={() => {
          const Members: Array<CID> = [];
          for (let i = 0; i < clients.length; i++) {
            if (checked[i]) {
              Members.push(clients[i].CID);
            }
          }
          send({ t: "MemberChoose", Members });
        }}
        disabled={checked.filter(id).length !== members}
      />
    </div>
  );
};
