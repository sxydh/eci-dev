package cn.net.bhe.eciportal.controller;

import cn.net.bhe.eciportal.model.R;
import cn.net.bhe.eciportal.model.dto.EciDTO;
import cn.net.bhe.eciportal.model.vo.EciVO;
import cn.net.bhe.eciportal.service.EciService;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/eci")
public class EciController {

    private final EciService eciService;

    public EciController(EciService eciService) {
        this.eciService = eciService;
    }

    @GetMapping("/list")
    public R<List<EciVO>> list() {
        return R.ok(eciService.queryEciList());
    }

    @PostMapping
    public R<EciVO> add(@RequestBody EciDTO eciDTO) {
        eciService.addEci(eciDTO);
        return R.ok();
    }

    @DeleteMapping
    public R<EciVO> delete(@RequestParam String name) {
        eciService.deleteEci(name);
        return R.ok();
    }

}
