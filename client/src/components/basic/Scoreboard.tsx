import React from "react";
import t8ns from "../../translations";
import { maxPts } from "../../shared";
import { Lang } from "../../etc";
import Checkbox from "./Checkbox";
import "./Scoreboard.css";

type Props = {
  lang: Lang;
  resPts: number;
  spyPts: number;
};

const points = Array.from(Array(maxPts), (_, i) => i + 1);

export default ({ lang, resPts, spyPts }: Props) => (
  <table className="Scoreboard">
    <tbody>
      <tr>
        <td>
          {points.map(k => (
            <Checkbox key={k} disabled checked={resPts >= k} />
          ))}
        </td>
        <td>
          <b>{t8ns[lang].resName}</b>
        </td>
      </tr>
    </tbody>
    <tr>
      <td>
        {points.map(k => (
          <Checkbox key={k} disabled checked={spyPts >= k} />
        ))}
      </td>
      <td>
        <b>{t8ns[lang].spyName}</b>
      </td>
    </tr>
  </table>
);
