

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart'as http;
import 'package:http/http.dart';
import 'package:task_manager2/all_task.dart';

class AddTaskPage extends StatelessWidget {
  //const AddTaskPage({Key? key}) : super(key: key);


  Future<Response> postData()async{
    Response response=await http.post(Uri.parse("http://192.168.0.204:8080/create?task=fldkfjlkj&task_detail=jjdgkgk")
      //   body: {
      // "task":"kfhksdkh",
      //     "task_detail":"hfjhjfhjdhjdfhh"
      //   },
        // headers: {
        //   "Content-Type":"application/json; charset=UTF-8"
        // }
    );

    return response;
  }


  TextEditingController nameController=TextEditingController();
  TextEditingController detailController=TextEditingController();


  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [


            Padding(
              padding: EdgeInsets.fromLTRB(30, 15, 30, 15),
              child:TextField(
                maxLines: 1,
                controller: nameController,
                decoration: InputDecoration(
                    border: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(30)
                    ),
                    filled: true,
                    fillColor: Colors.white70,
                    hintText: "Task name",
                    focusedBorder: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(10),
                        borderSide: BorderSide(
                            color: Colors.white,
                            width: 1.0
                        )
                    )
                ),
              ),),




            Padding(
              padding: EdgeInsets.fromLTRB(30, 15, 30, 15),
              child:TextField(
                maxLines: 5,
                controller: detailController,
                decoration: InputDecoration(
                    border: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(30)
                    ),
                    filled: true,
                    fillColor: Colors.white70,
                    hintText: "Task detail",
                    focusedBorder: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(10),
                        borderSide: BorderSide(
                            color: Colors.white,
                            width: 1.0
                        )
                    )
                ),
              ),
            ),


            ElevatedButton(onPressed: (){
              postData();

            }, child: Text("Add Task"))






          ],
        ),
      ),
    );
  }
}
