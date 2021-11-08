        // onload为浏览器对象中的事件,页面载入时触发 
        var sessionID;
        window.onload = function() {
            sessionID = randomString()
            var grid;
            var chessArr = [];
            var lineNum = 15;
            // 获取元素 
            var box = document.getElementById('chessboard');
            var chessBox = box.getElementsByTagName('div');
            var submitBtn = document.getElementById('submitBtn');
            submitBtn.onclick = function() {
                    var chessMaxNum = lineNum * lineNum;
                    var chessWH = 90 / lineNum;
                    for (var i = 0; i < chessMaxNum; i++) {
                        // 设置棋盘里小格子div元素 
                        grid = document.createElement('div');
                        grid.style.width = 'calc(' + chessWH + 'vmin - 2px)';
                        grid.style.height = 'calc(' + chessWH + 'vmin - 2px)';
                        grid.id = i;
                        box.appendChild(grid);
                        chessArr[i] = 0;
                        grid.onclick = function(x) {
                            // target 事件属性可返回事件的目标节点（触发该事件的节点），如生成事件的元素、文档或窗口。 
                            var index = x.target.id || x.target.parentNode.id;
                            return playChess(index);
                        };
                    };
                    mask.style.display = 'none';
                }
            //棋子对象 
            function Chess() {
                this.color = 'white';
                this.site = 0;
                // 创建一个class 
                this.chessDom = function() {
                    // 创造新节点 
                    var dom = document.createElement('p');
                    // 将这个名字给class 
                    dom.setAttribute('class', this.color);
                    return dom;
                }
            }

            function playChess(i) {
                if (chessArr[i] == 0) {
                    chessArr[i] = new Chess();
                    chessArr[i].color = 'white';
                    chessBox[i].appendChild(chessArr[i].chessDom());


                    // ai
                    console.log("ai's turn!")
                    aiCalculate(i)
                } else {
                    alert('此处有子!');
                }
            }

            function aiCalculate(i) {
                ajax_method("api", "index="+i+"&sessionID="+sessionID, "get", function(x){
                    arr = String(x).split(',')
                    if (x == "youWin"){
                        alert('you win')
                    } else if (arr[1] == "aiWin"){
                        console.log("success:", x)
                        chessArr[parseInt(x)] = new Chess();
                        chessArr[parseInt(x)].color = 'black';
                        chessBox[parseInt(x)].appendChild(chessArr[parseInt(x)].chessDom());
                        setTimeout(function(){
                            alert('ai win')
                         },0)
                        return 
                    }
                    console.log("success:", x)
                    chessArr[parseInt(x)] = new Chess();
                    chessArr[parseInt(x)].color = 'black';
                    chessBox[parseInt(x)].appendChild(chessArr[parseInt(x)].chessDom());
                })
            }

            function randomString(e) {
                var t = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
                a = t.length
                n = ""
                for (i = 0; i<32; i++) {
                    n +=t.charAt(Math.floor(Math.random()*a));
                }
                return n
            }
        };

