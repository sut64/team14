import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function Home() {
  const classes = useStyles();

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <h1 style={{ textAlign: "center" }}>ระบบบันทึกการเข้าใช้ห้อง</h1>
        <h4>Requirements</h4>
        <p>
        ระบบจัดการโรคติดต่อของโรงพยาบาล เป็นระบบที่มีไว้สำหรับเจ้าหน้าที่ในแผนกโรคติดต่อ เป็นระบบที่สามารถบันทึกข้อมูลเกี่ยวกับการใช้ห้องต่างๆของสถานพยาบาลไม่ว่าจะเป็นห้องพักผู้ป่วย หรือห้องที่ใช้ในการดำเนินงานต่างๆเช่นรักษาได้
พนักงานแผนกโรคติดต่อสามารถ login เข้าระบบเพื่อบันทึกข้อมูลการใช้ห้องต่างๆที่ต้องการจะใช้ผ่านระบบ ซึ่งพนักงานแผนกโรคติดต่อสามารถบันทึกข้อมูลต่างๆลงในระบบและบันทึกเป็นรายการการใช้ห้องได้


        </p>
      </Container>
    </div>
  );
}
export default Home;